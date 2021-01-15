//+build windows

package logger

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/influxdata/wlog"
	"golang.org/x/sys/windows/svc/eventlog"
)

const (
	LogTargetEventlog = "eventlog"
)

type eventLogger struct {
	logger *eventlog.Log
}

func (t *eventLogger) Write(b []byte) (n int, err error) {
	loc := prefixRegex.FindIndex(b)
	n = len(b)
	if loc == nil {
		err = t.logger.Info(1, string(b))
	} else if n > 2 { //skip empty log messages
		line := strings.Trim(string(b[loc[1]:]), " \t\r\n")
		switch rune(b[loc[0]]) {
		case 'I':
			err = t.logger.Info(1, line)
		case 'W':
			err = t.logger.Warning(2, line)
		case 'E':
			err = t.logger.Error(3, line)
		}
	}

	return
}

type eventLoggerCreator struct {
	logger *eventlog.Log
}

func (e *eventLoggerCreator) CreateLogger(config LogConfig) (io.Writer, error) {
	return wlog.NewWriter(&eventLogger{logger: e.logger}), nil
}

func RegisterEventLogger(name string) error {
	eventLog, err := eventlog.Open(name)
	if err != nil {
		log.Printf("E! An error occurred while initializing an event logger. %s", err)
		return err
	}

	registerLogger(LogTargetEventlog, &eventLoggerCreator{logger: eventLog})
	return nil
}
