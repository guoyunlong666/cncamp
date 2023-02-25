package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func header(w http.ResponseWriter, r *http.Request) {
	for k, valueSet := range r.Header {
		for _, value := range valueSet {
			w.Header().Add(k, value)
		}
	}
	w.Header().Set("Version", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)
	log.Println(r.RemoteAddr, http.StatusOK)
	io.WriteString(w, "ok")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

func main() {
	http.HandleFunc("/header", header)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
