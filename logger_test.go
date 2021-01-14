package elog

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	Errorf(Fields{}, "test1 %v", 1)
}
