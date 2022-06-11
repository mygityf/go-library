package log

import (
	"errors"
	"testing"
	"time"
)

func TestXlogImpl(t *testing.T) {
	impl := NewXlogImpl("../log/log-mock.log", DebugLevel)
	impl.Init()
	SetOutLogger(impl)
	endTime := time.Now().Add(time.Minute * 2)

	for time.Now().Before(endTime) {
		Debug("test debug").Any("p", 1).Any("q", 2).Line()
		Info("test Info").Any("p", 1).Line()
		Warn("test Warn").Any("p", 1).Line()
		Err("test Err").Error(errors.New("new err!")).Any("p", 1).Line()
		Fatal("test Fatal").Any("p", 1).Line()
		Debugf("test debugf-%v", 1).Any("p", 1).Any("q", 2).Line()
		Infof("test Infof-%v", 1).Any("p", 1).Line()
		Warnf("test Warnf-%v", 1).Any("p", 1).Line()
		Errorf("test Errorf-%v", 1).Error(errors.New("new err!")).Any("p", 1).Line()
		Fatalf("test Fatalf-%v", 1).Any("p", 1).Line()
		time.Sleep(time.Second)
	}
}
