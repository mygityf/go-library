package quit

import (
	"fmt"
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
	// counts active goroutines for GracefulStop
	serveWG sync.WaitGroup
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

// GracefulStop Graceful stop all running goroutines.
func (q *QuitEvent) GracefulStop() {
	q.Fire()
	q.WaitGoroutines()
}

// SignalHandler stop signal handle
func SignalHandler() {
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
