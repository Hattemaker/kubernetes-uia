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
	podName := os.Getenv("NAME")
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

		var fromName string
		if r.URL.Path != "/" && len(paths) >= 2 {
			fromName = paths[1]
		} else if podName != "" {
			fromName = podName
		} else {
			fromName = "App"
		}

		fmt.Fprintf(w, "Hei, Uia! Hilsen %s.\nTime %s\nVersion %s\n", fromName, time.Now().Local(), version)
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
