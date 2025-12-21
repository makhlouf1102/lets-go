package pkg

import "log"

type StandardLogger struct{}

func (l *StandardLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *StandardLogger) Verbose() bool {
	return true // Set to true to see every single step
}
