package main

import (
	"flag"
	"fmt"
	"github.com/CodeDing/http-transfer/config"
	"github.com/CodeDing/http-transfer/transfer"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
           Welcome to http tranfer service
        </p>
        </body>
</html>`
	fmt.Fprintln(w, html)
}

const (
	VERSION = "v1.1.3"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	version := fmt.Sprintf("config-%s-%d", VERSION, time.Now().Unix())
	fmt.Fprintf(w, version)
}

var (
	addr          string
	read_timeout  time.Duration
	write_timeout time.Duration
)

func init() {
	addr = fmt.Sprintf(":%s", config.String("service", "port"))
	read_timeout = time.Duration(config.IntDefault("service", "read_timeout", 300)) * time.Millisecond
	write_timeout = time.Duration(config.IntDefault("service", "write_timeout", 300)) * time.Millisecond
	//flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("logtostderr").Value.Set("false")
	flag.Lookup("v").Value.Set("5")
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logdir := filepath.Join(dir, "logs")
	flag.Lookup("log_dir").Value.Set(logdir)
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/version", versionHandler)

	mux.Handle("/api/v1/selector", &transfer.RouteHandler{})

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  read_timeout,
		WriteTimeout: write_timeout,
		Handler:      mux,
	}
	server.ListenAndServe()
}
