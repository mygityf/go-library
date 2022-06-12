package log

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

const (
	defaultRecordChanSize = 10000
	logTimeFormat         = `2006-01-02 15:04:05.000`
	logRotateFormat       = `2006010215`
	logRotateSuffixFormat = `20060102150405`
)

var (
	recordPool   *sync.Pool
	xlogImplPtr  *xlogImpl
	xlogImplOnce sync.Once
)

func init() {
	recordPool = &sync.Pool{New: func() interface{} {
		return &record{}
	}}
}

// one log line record
type record struct {
	// log time format string
	Timestamp string `json:"timestamp"`
	// log level
	LogLevel Level `json:"level"`
	// func code and line
	Caller string `json:"caller"`
	// log row
	Row string `json:"data"'`
}

// String log
func (r *record) String() []byte {
	buf := &bytes.Buffer{}
	buf.WriteString(r.Timestamp)
	buf.WriteString(" ")

	buf.WriteString(r.LogLevel.CapitalString())
	buf.WriteString(" ")

	buf.WriteString(r.Caller)
	buf.WriteString(" ")

	buf.WriteString(r.Row)
	buf.WriteString(" ")

	buf.WriteString("\n")
	return buf.Bytes()
}

// Json log
func (r *record) Json() []byte {
	res, _ := json.Marshal(r)
	return res
}

type xlogImpl struct {
	// log file name such as: log/xxx.log
	logName string
	// log level
	level Level
	// log record chain
	recordChan chan *record
	// inner file and buf writer
	file          *os.File
	fileBufWriter *bufio.Writer
	rotateFlag    string
}

// NewXlogImpl xlog impl
func NewXlogImpl(logName string, level Level) *xlogImpl {
	xlogImplOnce.Do(func() {
		xlogImplPtr = &xlogImpl{
			logName:    logName,
			level:      level,
			recordChan: make(chan *record, defaultRecordChanSize),
		}
		go xlogImplPtr.bootstrap()
	})
	return xlogImplPtr
}

// Init xlog
func (x *xlogImpl) Init() error {
	return x.createLogFile()
}

// Close close
func (x *xlogImpl) Close() {
	close(x.recordChan)
}

// Write to log chain
func (x *xlogImpl) Write(ctx context.Context, level Level, row string) {
	if level <= x.level {
		return
	}
	var (
		caller    string
		timestamp string
	)
	// source code, file and line num
	_, file, line, ok := runtime.Caller(2)
	if ok {
		caller = path.Base(file) + ":" + strconv.Itoa(line)
	}
	timestamp = time.Now().Format(logTimeFormat)
	r := recordPool.Get().(*record)
	r.Row = row
	r.Caller = caller
	r.LogLevel = level
	r.Timestamp = timestamp
	x.recordChan <- r
}

func (x *xlogImpl) createLogFile() error {
	if err := os.MkdirAll(path.Dir(x.logName), 0755); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	file, err := os.OpenFile(x.logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	x.file = file

	if x.fileBufWriter = bufio.NewWriterSize(x.file, 8192); x.fileBufWriter == nil {
		return errors.New("new fileBufWriter failed.")
	}

	return nil
}

// WriteToFile to file buf
func (x *xlogImpl) WriteToFile(r *record) error {
	if x.fileBufWriter == nil {
		return errors.New("no opened file")
	}
	if _, err := x.fileBufWriter.Write(r.String()); err != nil {
		return err
	}
	return nil
}

func (x *xlogImpl) rotate() error {
	logSuffix := time.Now().Format(logRotateFormat)
	if logSuffix == x.rotateFlag {
		return nil
	}

	if x.fileBufWriter != nil {
		if err := x.fileBufWriter.Flush(); err != nil {
			return err
		}
	}

	x.rotateFlag = logSuffix
	if x.file != nil {
		filePath := fmt.Sprintf("%v-%v", x.logName, time.Now().Format(logRotateSuffixFormat))
		if err := os.Rename(x.logName, filePath); err != nil {
			return err
		}
		if err := x.file.Close(); err != nil {
			return err
		}
	}

	return x.createLogFile()
}

func (x *xlogImpl) bootstrap() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("xlogImpl abort, unknown error, err:%v, stack:%s", err, string(debug.Stack()))
		}
	}()

	flushTimer := time.NewTimer(time.Millisecond * 500)
	rotateTimer := time.NewTimer(time.Minute)

	for {
		select {
		case r, ok := <-x.recordChan:
			if !ok {
				return
			}
			x.WriteToFile(r)
			recordPool.Put(r)
		case <-flushTimer.C:
			if x.fileBufWriter != nil {
				x.fileBufWriter.Flush()
			}
			flushTimer.Reset(time.Millisecond * 1000)
		case <-rotateTimer.C:
			x.rotate()
			rotateTimer.Reset(time.Minute)
		}
	}

}
