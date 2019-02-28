package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"gowork/xerr"
	"gowork/extern/logging"
	"os"
	"path/filepath"
	"runtime"
)

func initConfigure(base interface{}, v interface{}) error {
	path := flag.String("c", "conf/config.json", "-c config-file-path")
	flag.Parse()

	if v == nil {
		return ParseJSON(*path, base)
	}

	err := ParseJSON(*path, v)
	if err != nil {
		logging.Error("[initConfigure] parse config json failed, error = %s", err.Error())
		return err
	}

	txt, _ := json.Marshal(v)
	_ = json.Unmarshal(txt, base)

	return nil
}

func initLogger(appName, level, suffix string, daemon bool) error {
	// make up log file
	curPath := filepath.Dir(os.Args[0])
	absPath, _ := filepath.Abs(curPath)
	path := filepath.Join(absPath, "logging")
	path = filepath.Join(path, fmt.Sprintf("%s.log", appName))
	// create directory if not exist
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm); err != nil {
		return xerr.New(xerr.ERR_CODE_IO, "Failed to do create file: %s, error = %s", path, err.Error())
	}
	handler, err := logging.NewTimeRotationHandler(path, suffix, nil)
	if err != nil {
		return xerr.New(xerr.ERR_CODE_IO, "Failed to link file: %s, error = %s", path, err.Error())
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
