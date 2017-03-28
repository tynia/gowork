package application

import (
	"os"
	"fmt"
	"flag"
	"runtime"
	"path/filepath"
	"gowork/common/logging"
	e "gowork/common/error"
)

func initConfigure(v interface{}) *e.WError {
	path := flag.String("c", "conf/config.json", "-c config-file-path")
	flag.Parse()

	info, err := os.Stat(*path)
	if err != nil {
		if os.IsNotExist(err) {
			return e.NewWError(e.ERR_CODE_PARA, "Invalid config file: %s", *path)
		}
		return e.NewWError(e.ERR_CODE_PARA, "Failed to stat config file[path: %s]", *path)
	}
	
	mode := info.Mode()
	if mode.IsDir() {
		return e.NewWError(e.ERR_CODE_PARA, "Invalid config file[path: %s], it is a directory", *path)
	}

	if !mode.IsRegular() {
		return e.NewWError(e.ERR_CODE_PARA, "Invalid config file[path: %s], it is not a regular file", *path)
	}

	return parseJSON(*path, v)
}

func initLogger(appName, level, suffix string, daemon bool) *e.WError {
	// make up log file
	curPath := filepath.Dir(os.Args[0])
	absPath, _ := filepath.Abs(curPath)
	path := filepath.Join(absPath, "log")
	path = filepath.Join(path, fmt.Sprintf("%s.log", appName))
	// create directory if not exist
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm); err != nil {
		return e.NewWError(e.ERR_CODE_IO, "Failed to do create file: %s, error = %s", path, err.Error())
	}
	handler, err := logging.NewTimeRotationHandler(path, suffix, nil)
	if err != nil {
		return e.NewWError(e.ERR_CODE_IO, "Failed to link file: %s, error = %s", path, err.Error())
	}
	handler.SetLevelString(level)
	handler.SetFormat(func(appName, timeString string, rd *logging.Record) string {
		return "[" + timeString + "] " + " " + rd.Level.String() + " " + rd.Message + "\n"
	})
	logging.AddHandler(appName, handler)
	if daemon {
		logging.DisableStdout()
	}
	return nil
}

func initCPU(cpu int) {
	if cpu == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU()) //配0就用所有核
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}