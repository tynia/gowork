package service

import (
	"bytes"
	"encoding/json"
	e "gowork/error"
	"io"
	"io/ioutil"
	"os"
)

type configure struct {
	Log struct {
		Level  string // logging level (error/warning/info/debug)
		Suffix string // logging file name suffix
	}

	Prog struct {
		CPU        int // cpu in use
		Daemon     bool // you know
		HealthPort string // health port for monitor
	}

	Server struct {
		PortInfo string // serve port
	}
}

func ParseJSON(path string, v interface{}) *e.WError {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return e.NewWError(e.ERR_CODE_PARA, "Invalid config file: %s", path)
		}
		return e.NewWError(e.ERR_CODE_PARA, "Failed to stat config file[path: %s]", path)
	}

	mode := info.Mode()
	if mode.IsDir() {
		return e.NewWError(e.ERR_CODE_PARA, "Invalid config file[path: %s], it is a directory", path)
	}

	if !mode.IsRegular() {
		return e.NewWError(e.ERR_CODE_PARA, "Invalid config file[path: %s], it is not a regular file", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return e.NewWError(e.ERR_CODE_IO, "Failed to read config file[path: %s]", path)
	}
	var lines [][]byte
	buf := bytes.NewBuffer(data)
	for {
		line, err := buf.ReadBytes('\n')
		line = bytes.Trim(line, " \t\r\n")
		if len(line) > 0 && !bytes.HasPrefix(line, []byte("//")) {
			lines = append(lines, line)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return e.NewWError(e.ERR_CODE_IO, "Failed to read config file[path: %s] content", path)
		}
	}

	data = bytes.Join(lines, []byte{})
	err = json.Unmarshal(data, v)
	if err != nil {
		return e.NewWError(e.ERR_CODE_IO, "Failed to unmarshal file[path: %s] content to json", path)
	}

	return nil
}
