package logger

import (
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogrusLogger() {
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ", lineNumber:" + strconv.Itoa(frame.Line)
			return frame.Function, fileName
		},
	})
	Log.SetLevel(logrus.InfoLevel)

}
