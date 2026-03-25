package main

import (
	"database/sql"
	"fmt"
	"net/http"

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

	// Vulnerable endpoint for CodeQL testing — SQL injection from user input
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		db, _ := sql.Open("sqlite3", ":memory:")
		rows, _ := db.Query("SELECT * FROM items WHERE name = '" + query + "'")
		defer rows.Close()
		fmt.Fprint(w, "searched")
	})

	http.ListenAndServe(":80", nil)
}
