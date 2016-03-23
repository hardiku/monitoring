package log

import (
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

//Log is an exported logging variable
var Log = logging.MustGetLogger("example")

func init() {
	//	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	//logging.SetBackend(backend1Leveled, backend2Formatter)
	logging.SetBackend(backend2Formatter)

	Log.Info("Initialized logging.")
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	Log.Debug(args...)
}
func Info(args ...interface{}) {
	Log.Info(args...)
}
func Notice(args ...interface{}) {
	Log.Notice(args...)
}
func Warning(args ...interface{}) {
	Log.Warning(args...)
}
func Error(args ...interface{}) {
	Log.Error(args...)
}
func Critical(args ...interface{}) {
	Log.Critical(args...)
}
func Panic(args ...interface{}) {
	Log.Panic(args...)
}
