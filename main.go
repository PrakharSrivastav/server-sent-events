package main

import (
	"fmt"
	"github.com/tinygg/gofaker"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// route to html
	http.Handle("/", http.FileServer(http.Dir("static")))

	// route to sse
	http.HandleFunc("/events", eventsHandler)

	// serve application
	http.ListenAndServe(":8080", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	count := 1
	gofaker.Seed(int64(rand.Int()))

	for {
		if count > 10 {
			break
		}

		breed := gofaker.Cat()
		fmt.Fprintf(w, "data: %s \n\n", fmt.Sprintf("%d : %s", count, breed))
		w.(http.Flusher).Flush()

		time.Sleep(100 * time.Millisecond)

		count++
	}

}
