package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Gopher! 👋")
	if r.TLS != nil && r.TLS.HandshakeComplete {
		fmt.Fprintln(w, "You are connected securely! ✅")
	} else {
		fmt.Fprintln(w, "You are connected insecurely! 💥")
	}
}

func main() {
	http.HandleFunc("/", hello)

	var err error

	// Randomly choose between using TLS and not using TLS.
	switch rand.Intn(2) {
	case 0:
		fmt.Fprintf(os.Stderr, "listening at :8080 using TLS\n")
		err = http.ListenAndServeTLS(":8080", "localhost.pem", "localhost-key.pem", nil)
	case 1:
		fmt.Fprintf(os.Stderr, "listening at :8080 without TLS\n")
		err = http.ListenAndServe(":8080", nil)
	}

	log.Fatal(err)
}
