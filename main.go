package main

import (
	"fmt"
	"github.com/tinygg/gofaker"
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
	gofaker.Seed(int64(rand.Int()))

	for {
		if count > 10 {
			// break out of the loop on the 11th iteration, to drop the http connection
			break
		}

		breed := gofaker.Cat()
		// sleep for 100ms to simulate server processing
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(w, "data: %s \n\n", fmt.Sprintf("%d : %s", count, breed))

		// flush the response
		w.(http.Flusher).Flush()

		count++
	}
}
