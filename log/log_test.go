package log

import (
	"errors"
	"testing"
)

func TestLogChain_Any(t *testing.T) {
	Debug("test debug").Any("p", 1).Any("q", 2).Line()
	Info("test Info").Any("p", 1).Line()
	Warn("test Warn").Any("p", 1).Line()
	Err("test Err").Error(errors.New("new err!")).Any("p", 1).Line()
	Fatal("test Fatal").Any("p", 1).Line()
}
