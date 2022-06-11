package log

import "context"

// OutLogger logger interface
type OutLogger interface {
	Write(ctx context.Context, level Level, row string)
	Close()
}

var (
	defaultOutLogger OutLogger
)

// SetOutLogger set logger
func SetOutLogger(logger OutLogger) {
	defaultOutLogger = logger
}
