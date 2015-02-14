package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/juju/loggo"
	"github.com/zbindenren/negroni-loggo"
)

// custom loggo formatter
type logFormatter struct{}

func (*logFormatter) Format(level loggo.Level, module, filename string, line int, timestamp time.Time, message string) string {
	ts := timestamp.In(time.UTC).Format("2006-01-02 15:04:05")
	filename = filepath.Base(filename)
	return fmt.Sprintf("%s %s [%s] %s", ts, level, module, message)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success!\n")
	})

	n := negroni.New()
	// use the custom loggo formatter
	n.Use(negroniloggo.NewLoggerWithCustomWriter(loggo.NewSimpleWriter(os.Stderr, &logFormatter{})))
	// to use the default loggo formatter
	// n.Use(negronilogging.NewLogger())
	n.UseHandler(r)

	n.Run(":3000")
}
