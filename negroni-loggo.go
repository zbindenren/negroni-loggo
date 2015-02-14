package negroniloggo

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/juju/loggo"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	*loggo.Logger
}

// NewLogger returns a new *Logger with the standard loggo writer.
func NewLogger(moduleName string) *Logger {
	log := loggo.GetLogger(moduleName)
	loggo.ConfigureLoggers(moduleName + "=INFO")
	return &Logger{&log}
}

// NewLoggerWithCustomWriter returns a new *Logger with with a custom writer.
func NewLoggerWithCustomWriter(moduleName string, writer loggo.Writer) *Logger {
	log := NewLogger(moduleName)
	loggo.ReplaceDefaultWriter(writer)
	return log
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	l.Infof("started %s %s", r.Method, r.URL.Path)

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	l.Infof("completed %v %s in %v", res.Status(), http.StatusText(res.Status()), time.Since(start))
}
