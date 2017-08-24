// +build linux darwin dragonfly freebsd netbsd openbsd solaris

package gsyslog

import (
	"fmt"
	"log/syslog"
)

// builtinLogger wraps the Golang implementation of a
// syslog.Writer to provide the Syslogger interface
type builtinLogger struct {
	*builtinWriter
}

// NewLogger is used to construct a new Syslogger
func NewLogger(p Priority, tag string) (Syslogger, error) {
	l, err := newBuiltin(syslog.Priority(p), tag)
	if err != nil {
		return nil, err
	}
	return &builtinLogger{l}, nil
}

// DialLogger is used to construct a new Syslogger that establishes connection to remote syslog server
func DialLogger(network, raddr string, p Priority, tag string) (Syslogger, error) {
	l, err := dialBuiltin(network, raddr, syslog.Priority(p), tag)
	if err != nil {
		return nil, err
	}

	return &builtinLogger{l}, nil
}

// WriteLevel writes out a message at the given priority
func (b *builtinLogger) WriteLevel(p Priority, buf []byte) error {
	var err error
	m := string(buf)
	switch p {
	case LOG_EMERG:
		_, err = b.writeAndRetry(syslog.LOG_EMERG, m)
	case LOG_ALERT:
		_, err = b.writeAndRetry(syslog.LOG_ALERT, m)
	case LOG_CRIT:
		_, err = b.writeAndRetry(syslog.LOG_CRIT, m)
	case LOG_ERR:
		_, err = b.writeAndRetry(syslog.LOG_ERR, m)
	case LOG_WARNING:
		_, err = b.writeAndRetry(syslog.LOG_WARNING, m)
	case LOG_NOTICE:
		_, err = b.writeAndRetry(syslog.LOG_NOTICE, m)
	case LOG_INFO:
		_, err = b.writeAndRetry(syslog.LOG_INFO, m)
	case LOG_DEBUG:
		_, err = b.writeAndRetry(syslog.LOG_DEBUG, m)
	default:
		err = fmt.Errorf("Unknown priority: %v", p)
	}
	return err
}
