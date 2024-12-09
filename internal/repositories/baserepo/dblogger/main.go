package dblogger

import (
	"log"
)

type LoggingEventReceiver struct{}

func (l *LoggingEventReceiver) Event(eventName string) {}
func (l *LoggingEventReceiver) EventKv(eventName string, kvs map[string]string) {
	if eventName == "dbr.query" {
		// kvs["sql"] contém a query
		// kvs["args"] contém os argumentos
		log.Printf("[QUERY] sql=%s args=%s", kvs["sql"], kvs["args"])
	}
}
func (l *LoggingEventReceiver) EventErr(eventName string, err error) error {
	log.Printf("[ERROR] %s: %v", eventName, err)
	return err
}
func (l *LoggingEventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	log.Printf("[ERROR] %s: %v | %v", eventName, err, kvs)
	return err
}
func (l *LoggingEventReceiver) Timing(eventName string, nanoseconds int64)                          {}
func (l *LoggingEventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {}
