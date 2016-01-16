package logger

import (
	"io"
	"log/syslog"
	"os"
	"path/filepath"
)

var (
	Facility = syslog.LOG_DAEMON
	Severity = syslog.LOG_INFO
	Prefix   = filepath.Base(os.Args[0])
)

func New(priority syslog.Priority, prefix string) (io.Writer, error) {
	return syslog.New(priority, prefix)
}

func Priority(severity string) syslog.Priority {
	var priority syslog.Priority
	switch severity {
	case "error":
		priority = syslog.LOG_ERR | Facility
	case "warning":
		priority = syslog.LOG_WARNING | Facility
	case "debug":
		priority = syslog.LOG_DEBUG | Facility
	default:
		priority = syslog.LOG_INFO | Facility
	}
	return priority
}

func Error(msg string) {
	w, _ := syslog.New(Facility, Prefix)
	w.Err(msg)
}

func Fatal(msg string) {
	w, _ := syslog.New(Facility, Prefix)
	w.Crit(msg)
	os.Exit(1)
}
