package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var i int

func main() {
	// Hello world, the web server

	ticker := time.NewTicker(1 * time.Second)
	//quit := make(chan struct{})
	go readPhase(ticker)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, fmt.Sprintf("Hello, world! %d\n", i))
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func readPhase(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			i++
			resp, err := http.Get("http://localhost:8000/hello")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			for i := 0; scanner.Scan() && i < 5; i++ {
				fmt.Println(scanner.Text())
			}
		}
	}
}
