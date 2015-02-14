package negronilogging

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/juju/loggo"
)

const (
	ModuleName = "request"
)

// LoggoLogerMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	*loggo.Logger
}

// NewMiddleware returns a new *Middleware, yay!
func NewLogger() *Logger {
	log := loggo.GetLogger(ModuleName)
	loggo.ConfigureLoggers(ModuleName + "=INFO")
	return &Logger{&log}
}

func NewLoggerWithCustomWriter(writer loggo.Writer) *Logger {
	log := loggo.GetLogger(ModuleName)
	loggo.ConfigureLoggers(ModuleName + "=INFO")
	loggo.ReplaceDefaultWriter(writer)
	return &Logger{&log}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	l.Infof("started %s %s", r.Method, r.URL.Path)

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	l.Infof("completed %v %s in %v", res.Status(), http.StatusText(res.Status()), time.Since(start))
}
