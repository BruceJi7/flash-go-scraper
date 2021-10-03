package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var print = fmt.Println

func logError(err *error) {
	if *err != nil {
		log.Fatal(*err)
	}
}

func handleError(err *error, message string) {
	if *err != nil {
		log.Fatal(*err)
	} else {
		print(message)
	}
}

const REQUESTURL = "https://endic.naver.com/search.nhn?sLn=en&searchOption=all&query="

func main() {
	print("Flashscraper")
	print(REQUESTURL)

	//Setup client, use timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Initialise HTTP request before sending it
	url := REQUESTURL + "help"
	request, err := http.NewRequest("GET", url, nil)
	handleError(&err, "Request constructed")

	request.Header.Set("User-Agent", "Chrome/85.0.4183.12")

	response, err := client.Do(request)
	handleError(&err, "Response received")
	defer response.Body.Close()

}
