package quit

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	gQuitEvent *QuitEvent
	once       sync.Once
)

// init
func init() {
	gQuitEvent = NewQuitEvent()
}

// GetQuitEvent get sigleton quit event
func GetQuitEvent() *QuitEvent {
	once.Do(func() {
		if gQuitEvent == nil {
			gQuitEvent = NewQuitEvent()
		}
	})
	return gQuitEvent
}

// QuitEvent quit event struct
type QuitEvent struct {
	*Event
	// quit closer list to be close
	quitCloserList []QuitCloser
	// io closer list to be close
	closerList []io.Closer
	// stop func list
	stopFuncList []func()
	// counts active goroutines for GracefulStop
	serveWG sync.WaitGroup
}

// QuitCloser Shutdown
type QuitCloser interface {
	// Once Shutdown has been called on a server, it may not be reused;
	// future calls to methods such as Serve will return ErrServerClosed.
	Shutdown(ctx context.Context) error
}

// NewQuitEvent returns a new, ready-to-use Event.
func NewQuitEvent() *QuitEvent {
	return &QuitEvent{
		Event: NewEvent(),
	}
}

// AddGoroutine Incr count of running goroutine
func (q *QuitEvent) AddGoroutine() {
	q.serveWG.Add(1)
}

// DoneGoroutine Decr count of running goroutine
func (q *QuitEvent) DoneGoroutine() {
	q.serveWG.Done()
}

// WaitGoroutines Waiting all running goroutine quit.
func (q *QuitEvent) WaitGoroutines() {
	q.serveWG.Wait()
}

// AddQuitCloser closer will be called before goroutine quit.
func (q *QuitEvent) AddQuitCloser(closer QuitCloser) {
	q.quitCloserList = append(q.quitCloserList, closer)
}

// AddCloser closer will be called before goroutine quit.
func (q *QuitEvent) AddCloser(closer io.Closer) {
	q.closerList = append(q.closerList, closer)
}

// AddStopFunc stop func will be called before goroutine quit.
func (q *QuitEvent) AddStopFunc(stopFunc func()) {
	q.stopFuncList = append(q.stopFuncList, stopFunc)
}

// GracefulStop Graceful stop all running goroutines.
func (q *QuitEvent) GracefulStop() {
	q.Fire()
	for _, closer := range q.quitCloserList {
		if closer != nil {
			_ = closer.Shutdown(context.TODO())
		}
	}
	for _, closer := range q.closerList {
		if closer != nil {
			_ = closer.Close()
		}
	}
	for _, stopFunc := range q.stopFuncList {
		if stopFunc == nil {
			stopFunc()
		}
	}
	q.WaitGoroutines()
}

// WaitSignal stop signal handle
func WaitSignal() {
	shutdownHook := make(chan os.Signal, 1)
	signal.Notify(shutdownHook,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	localPid := os.Getpid()
	sig := <-shutdownHook

	fmt.Printf("caught sig exit sig:%v,localPid:%v", sig, localPid)
	go func() {
		GetQuitEvent().GracefulStop()
	}()
	// wait 3 second for quit event graceful stop.
	time.Sleep(3 * time.Second)
	os.Exit(0)
}
