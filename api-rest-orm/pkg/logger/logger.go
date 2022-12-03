package logger

/** Desacoplar y simplificar la libreria Logrus Logs */
import (
	"github.com/sirupsen/logrus"
	"os"
)

const eventName = "x_event"

type Logger interface {
	Trace(msg string)
	Info(msg string)
	Debug(msg string)
	Error(msg string)
	Fatal(msg string)
	Warn(msg string)
	TraceF(msg string, logFields interface{})
	InfoF(msg string, logFields interface{})
	Fields(msg string, fields logrus.Fields)
	DebugF(msg string, logFields interface{})
	ErrorF(msg string, logFields interface{})
	FatalF(msg string, logFields interface{})
	WarnF(msg string, logFields interface{})
}

type loggerLogrus struct {
	logger *logrus.Logger
	level  string
}

func NewLogger(level string) Logger {
	log := &loggerLogrus{level: level}
	log.configLogs()
	return log
}

func (l *loggerLogrus) configLogs() {
	l.logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{PrettyPrint: false},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	l.setLogLevel(l.level)
}

func (l *loggerLogrus) setLogLevel(logLevel string) {
	if level, error := logrus.ParseLevel(logLevel); error != nil {
		l.logger.Level = logrus.TraceLevel
	} else {
		l.logger.Level = level
	}
}

func (l *loggerLogrus) Trace(msg string) {
	l.logger.Trace(msg)
}
func (l *loggerLogrus) Info(msg string) {
	l.logger.Info(msg)
}
func (l *loggerLogrus) Debug(msg string) {
	l.logger.Debug(msg)
}
func (l *loggerLogrus) Error(msg string) {
	l.logger.Error(msg)
}
func (l *loggerLogrus) Fatal(msg string) {
	l.logger.Fatal(msg)
}
func (l *loggerLogrus) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *loggerLogrus) TraceF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Trace(msg)
}
func (l *loggerLogrus) InfoF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Info(msg)
}
func (l *loggerLogrus) Fields(msg string, fields logrus.Fields) {
	l.logger.WithFields(fields).Info(msg)
}
func (l *loggerLogrus) DebugF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Debug(msg)
}
func (l *loggerLogrus) ErrorF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Error(msg)
}
func (l *loggerLogrus) FatalF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Fatal(msg)
}
func (l *loggerLogrus) WarnF(msg string, logFields interface{}) {
	l.logger.WithFields(logrus.Fields{eventName: logFields}).Warn(msg)
}
