package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	serviceURL := os.Getenv("GREETER_SERVICE_URL")

	for {
		time.Sleep(5 * time.Second)

		// Ex: http://localhost:8080/Repeater
		resp, err := http.DefaultClient.Get(fmt.Sprintf("%s/Repeater", serviceURL))
		if err != nil {
			log.Println(err)
			continue
		}

		body, err := readResponse(resp)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(body)
	}
}

func readResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
