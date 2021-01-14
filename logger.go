package elog

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

type Fields = logrus.Fields

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
}

func Errorf(fields Fields, fmt string, args ...interface{}) {
	log.WithFields(fields).Errorf(fmt, args...)
}
