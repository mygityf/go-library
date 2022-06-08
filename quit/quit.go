package quit

import "sync"

var (
	gQuitEvent *QuitEvent
	once       sync.Once
)

// init
func init() {
	gQuitEvent = NewQuitEvent()
}

// GetQuitEvent 获取一个安全退出处理单例
func GetQuitEvent() *QuitEvent {
	once.Do(func() {
		if gQuitEvent == nil {
			gQuitEvent = NewQuitEvent()
		}
	})
	return gQuitEvent
}

// QuitEvent 安全退出处理器
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
