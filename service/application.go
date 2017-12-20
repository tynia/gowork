package service

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"gowork/extern/logging"
	"gowork/net"
	e "gowork/error"
)

var (
	shutdown   = make(chan os.Signal)
	running    = make(chan bool)
	hup        = make(chan os.Signal)
	dumpPrefix = "panic"
	dumpMode   = os.FileMode(0777)
	dumpFlag   = os.O_CREATE | os.O_WRONLY
)

func reviewDumpPanic(file *os.File) *e.WError {
	fileinfo, err := file.Stat()
	if err != nil {
		return e.WrapError(e.ERR_CODE_SYS, err)
	}
	if fileinfo.Size() == 0 {
		file.Close()
		err := os.Remove(file.Name())
		if err != nil {
			return e.WrapError(e.ERR_CODE_IO, err)
		}
	}
	return nil
}

type InitHandlerFunc func() *e.WError
type Handler net.Handler

var (
	App *Application
)

type Application struct {
	appName     string
	baseConfig  *configure
	health      Handler
	handler     Handler
	initHandler InitHandlerFunc
	config      map[string]string
	handlerMap  map[string]net.HandlerFunc
}

func NewApplication(name string, conf map[string]string) *Application {
	App = &Application{
		appName:     name,
		baseConfig:  new(configure),
		health:      nil,
		handler:     nil,
		initHandler: nil,
		config:      make(map[string]string),
		handlerMap:  make(map[string]net.HandlerFunc),
	}

	for k, v := range conf {
		App.Set(k, v)
	}

	return App
}

func (app *Application) signal() {
	signal.Notify(shutdown, syscall.SIGINT)
	signal.Notify(shutdown, syscall.SIGTERM)
	signal.Notify(hup, syscall.SIGHUP)
	go func() {
		for {
			select {
			case <-shutdown:
				logging.Info("[Application.signal] receive signal SIGINT or SIGTERM, to stop server...")
				running <- false
			case <-hup:
			}
		}
	}()
	logging.Info("[Application.signal] register signal ok")
}

func (app *Application) dump() (*os.File, *e.WError) {
	suffix := fmt.Sprintf("-dump-%s", app.appName)
	filename := dumpPrefix + suffix + "." + strconv.Itoa(os.Getpid())
	file, err := os.OpenFile(filename, dumpFlag, dumpMode)
	if err != nil {
		return file, e.WrapError(e.ERR_CODE_SYS, err)
	}

	if err := syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		return file, e.WrapError(e.ERR_CODE_SYS, err)
	}
	return file, nil
}

func (app *Application) initBaseModule() *e.WError {

	err := initLogger(app.appName,
		app.baseConfig.Log.Level,
		app.baseConfig.Log.Suffix,
		app.baseConfig.Prog.Daemon)
	if err != nil {
		logging.Error("[Application.initBaseModule] Initialize logger moduler failed, error = %s", err.Error())
		return err
	}

	initCPU(app.baseConfig.Prog.CPU)

	return nil
}

func (app *Application) register() {
	for k, v := range app.handlerMap {
		net.HandleFunc(k, v)
		logging.Info("[Application.register] register handler[addr: %s]", k)
	}
}

func (app *Application) run() *e.WError {
	// service
	if app.baseConfig.Server.PortInfo == "" {
		logging.Warning("[application.run] not valid serve port, try to run with no serve")
		return nil //e.NewWError(e.ERR_CODE_PARA, "Invalid Serve port for application, port: %#+v", app.baseConfig.Server.PortInfo)
	} else {
		net.Serve(app.baseConfig.Server.PortInfo, app.handler)
	}

	// health
	if app.baseConfig.Prog.HealthPort != "" && app.health != nil {
		net.Health(app.baseConfig.Prog.HealthPort, app.health)
	}

	return nil
}

func (app *Application) SetHealthHandler(handler Handler) {
	app.health = handler
}

func (app *Application) SetServeHandler(handler Handler) {
	app.handler = handler
}

func (app *Application) SetInitHandler(handler InitHandlerFunc) {
	app.initHandler = handler
}

func (app *Application) Set(key string, value string) {
	if v, ok := app.config[key]; ok {
		logging.Warning("[Application.Set] Try to replace value[%#+v] to key = %s, original value: %s", value, key, v)
	}

	app.config[key] = value
	logging.Info("[Application.Set] Add/Replace [key: %s, value: %#+v] into config ok", key, value)
}

func (app *Application) Get(key string) string {
	if v, ok := app.config[key]; ok {
		return v
	}

	logging.Error("[Application.Get] Failed to get value of key[%s], value is NULL", key)
	return ""
}

func (app *Application) AddHandlerFunc(addr string, handler net.HandlerFunc) {
	_, ok := app.handlerMap[addr]
	if ok {
		logging.Warning("[Application.AddHandlerFunc] Try to replace handler to addr = %s", addr)
	}

	app.handlerMap[addr] = handler
	logging.Info("[Application.AddHandlerFunc] Add/Replace [addr: %s] ok", addr)
}

func (app *Application) Go() *e.WError {
	logging.Info("[Application] start")
	// register signal
	app.signal()

	// dump when error occurs
	file, err := app.dump()
	if err != nil {
		logging.Error("[Application.Go] Error occurs when initialize dump panic file, error = %s", err.Error())
	}

	// output exit info
	defer func() {
		logging.Info("[Application.Go] server stop...code: %d", runtime.NumGoroutine())
		time.Sleep(time.Second)
		logging.Info("[Application.Go] server stop...ok")
		if err := reviewDumpPanic(file); err != nil {
			logging.Error("[Application.Go] Failed to review dump panic file, error = %s", err.Error())
		}
	}()

	// parse config file content
	total := map[string]interface{}{}
	err = initConfigure(app.baseConfig, &total)
	if err != nil {
		logging.Error("[Application.Go] Cannot parse config file, error = %s", err.Error())
		return err
	}

	for k, v := range total {
		if str, ok := v.(string); ok {
			app.Set(k, str)
		}
	}

	// init logger
	err = app.initBaseModule()
	if err != nil {
		logging.Error("[Application.Go] Cannot init logger module, error = %s", err.Error())
		return err
	}

	if app.initHandler != nil {
		err := app.initHandler()
		if err != nil {
			logging.Error("[Application.Go] Error occurs when initialize application, error = %s", err.Error())
			return err
		}
	}

	// register handler
	app.register()
	// run
	err = app.run()
	if err != nil {
		logging.Error("[Application.Go] Error in running application, error = %s", err.Error())
		return err
	}

	<-running
	return err
}
