package log

import (
	"encoding/json"
	"io"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) String() string {
	return [...]string{"DEBUG", "INFO ", "WARN ", "ERROR", "FATAL", "OFF"}[l]
}

type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) PrintDebug(message string, properties map[string]string) {
	l.print(LevelDebug, message, properties)
}

func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

func (l *Logger) PrintWarn(message string, properties map[string]string) {
	l.print(LevelWarn, message, properties)
}

func (l *Logger) PrintError(message string, properties map[string]string) {
	l.print(LevelError, message, properties)
}

func (l *Logger) PrintFatal(message string, properties map[string]string) {
	l.print(LevelFatal, message, properties)
}

func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	if level == LevelDebug || level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	var line []byte

	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message: " + err.Error())
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, string(message), nil)
}
