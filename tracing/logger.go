package tracing

import "github.com/eztop/go_common/log"

type jaegerLogger struct{}

func (l *jaegerLogger) Error(msg string) {
	log.Errorf(msg)
}

func (l *jaegerLogger) Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func (l *jaegerLogger) Debugf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}
