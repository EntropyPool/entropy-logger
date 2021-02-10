package elog

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	Errorf(Fields{}, "test1 %v", 1)
}

func TestInfof(t *testing.T) {
	Infof(Fields{}, "test1 %v", 2)
}

func TestDebugf(t *testing.T) {
	Debugf(Fields{}, "test1 %v", 3)
}

func TestSetLevel(t *testing.T) {
	SetLevel("panic")
	SetLevel("fatal")
	SetLevel("error")
	SetLevel("warn")
	SetLevel("info")
	SetLevel("debug")
	SetLevel("trace")
}
