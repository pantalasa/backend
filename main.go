package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	quotes, err := Quotes()
	if err != nil {
		logrus.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "15")
		quote := RandomQuote(quotes)

		logrus.Info(quote)
		fmt.Fprint(w, quote)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		logrus.Debug("healthz probe")
		fmt.Fprint(w, "ok")
	})

	// Configure explicit timeouts so the server isn't vulnerable to slow-loris
	// clients that can otherwise hold connections open indefinitely.
	srv := &http.Server{
		Addr:              ":80",
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}

// healthz endpoint for k8s probes
