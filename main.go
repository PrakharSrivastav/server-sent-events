package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// serve html content on the path '/'
	http.Handle("/", http.FileServer(http.Dir("static")))

	// serve sse on the path '/events'
	http.HandleFunc("/events", eventsHandler)

	// run http server on port 8080
	http.ListenAndServe(":8080", nil)
}

// eventsHandler receives the request, sets up the correct response headers
// and returns 10 cat breeds to the client.
// On 11th iteration, the control breaks out of the loop and the
// http connection is dropped.
// This connection drop is captured by client and client handles the connection drop accordingly

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// set up the required response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// content-type should be 'text/event-stream'
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	// keep the connection alive for this response
	w.Header().Set("Connection", "keep-alive")

	count := 1
	gofakeit.Seed(int64(rand.Int()))

	numIterations := 10

	for count <= numIterations {
		breed := gofakeit.Cat()
		// sleep for 100ms to simulate server processing
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(w, "data: %d : %s \n\n", count, breed)

		// check if the ResponseWriter supports Flusher interface
		if f, ok := w.(http.Flusher); ok {
			// flush the response
			f.Flush()
		}

		count++
	}
}
