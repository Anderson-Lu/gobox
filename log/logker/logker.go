package log_helper

import "os"
import logging "github.com/op/go-logging"

var logker *logging.Logger

func initLogger() {
	logker = logging.MustGetLogger("example")
	var format = logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{level:.4s} ▶ %{shortfunc} %{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

func NewLogger() *logging.Logger {
	if logker == nil {
		initLogger()
	}
	return logker
}
