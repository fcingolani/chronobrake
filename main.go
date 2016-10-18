package main

import (
	"flag"
	"io"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

var (
	port string
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	flag.StringVar(&port, "port", "3000", "server port")
	flag.Parse()

	http.HandleFunc("/", ChronoBrakeServer)

	log.Info("Listening on ", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func ChronoBrakeServer(w http.ResponseWriter, r *http.Request) {

	log.WithFields(log.Fields{
		"method":  r.Method,
		"host":    r.Host,
		"url":     r.URL,
		"headers": r.Header,
	}).Info()

	timeStr := r.URL.Path[1:len(r.URL.Path)]
	timeInt, err := strconv.Atoi(timeStr)

	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 400)
		return
	}

	time.Sleep(time.Duration(timeInt) * time.Millisecond)

	io.WriteString(w, timeStr)

}
