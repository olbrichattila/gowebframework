package logger

import (
	"encoding/json"
	"framework/internal/app/storage"
	"time"
)

const (
	typeInfo     = "info"
	typeWarning  = "warning"
	typeError    = "error"
	typeCritical = "critical"
)

func New(storage storage.Storager) Logger {
	return &Logging{
		storage:  storage,
		filename: "log/app.log", // TODO make it configurable from conf.
	}
}

type Logger interface {
	Info(string)
	Warning(string)
	Error(string)
	Critical(string)
}

type Logging struct {
	storage  storage.Storager
	filename string
}

func (l *Logging) Info(message string) {
	l.log(typeInfo, message)
}

func (l *Logging) Warning(message string) {
	l.log(typeWarning, message)
}

func (l *Logging) Error(message string) {
	l.log(typeError, message)
}

func (l *Logging) Critical(message string) {
	l.log(typeCritical, message)
}

func (l *Logging) log(ltype, message string) {

	logStr := map[string]string{
		"type":       ltype,
		"message":    message,
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	logLine, err := json.Marshal(logStr)
	if err != nil {
		return
	}

	l.storage.Append(l.filename, string(logLine)+"\n")
}
