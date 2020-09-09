package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	name := os.Getenv("NAME")
	version := os.Getenv("VERSION")

	log.Printf("Application version %s\n", version)
	log.Println("Starting application")

	// Denne er kun for å demonstrere hvordan Kubernetes håndterer applikasjoner
	// som ikke klarer å starte.
	if version == "3" {
		log.Println("Waiting")
		time.Sleep(5 * time.Second)
		log.Fatal("Failed to start application")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")

    from:= r.UserAgent()
		if r.URL.Path != "/" && len(paths) >= 2 {
			from= paths[1]
		}

		if name == "" {
		  name = "App"
		}

		fmt.Fprintf(w, "Jeg er %s og kjører i versjon %s\n", name, version)
		fmt.Fprintf(w, "Jeg har fått melding fra %s\n", from)
		fmt.Fprintf(w, "Klokken er: %s\n", time.Now().Local())
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
