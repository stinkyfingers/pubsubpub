package main

import (
	"github.com/stinkyfingers/pubsubpub/pubsub"

	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "Listening Port --port=8080")
)

func main() {

	flag.Parse()
	if port == nil {
		log.Fatal("failed to setup listening port")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", status)
	mux.HandleFunc("/pub", publish)

	log.Printf("Starting server on %s", *port)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), mux)
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Status: OK")
	return
}

func publish(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	topic := r.URL.Query().Get("topic")
	pubsub.Push(topic, data)
	return
}
