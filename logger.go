package elog

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var log = logrus.New()

type Fields = logrus.Fields

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
}

func logEntry() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return log.WithField("fatal", "cannot get caller")
	}

	paths := strings.Split(file, "/")
	var filePath string

	j := 0
	i := len(paths) - 1

	for j < 2 && 0 <= i {
		if 0 < len(filePath) {
			filePath = paths[i] + "/" + filePath
		} else {
			filePath = paths[i]
		}
		i -= 1
		j += 1
	}

	filename := filePath + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return log.WithField("file", filename).WithField("func", fn)
}

func Errorf(fields Fields, fmt string, args ...interface{}) {
	logEntry().WithFields(fields).Errorf(fmt, args...)
}

func Fatalf(fields Fields, fmt string, args ...interface{}) {
	logEntry().WithFields(fields).Fatalf(fmt, args...)
}

func Infof(fields Fields, fmt string, args ...interface{}) {
	logEntry().WithFields(fields).Infof(fmt, args...)
}

func Debugf(fields Fields, fmt string, args ...interface{}) {
	logEntry().WithFields(fields).Debugf(fmt, args...)
}

func SetLevel(levelStr string) error {
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
