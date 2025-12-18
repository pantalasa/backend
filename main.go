package main

import (
	"fmt"
	"net/http"

	"github.com/ProtonMail/go-crypto/openpgp" // GPL-licensed library (example for SBOM policy)
	"github.com/sirupsen/logrus"
)

func main() {
	// Example usage of GPL-licensed library (for SBOM policy demonstration)
	_ = openpgp.EntityList{}

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

		// fmt.Fprint(w, "Hello from Earthly - 12:28PM")

	})

	http.ListenAndServe(":80", nil)
}
