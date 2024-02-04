package main

import (
	"github.com/therealdannzor/go-ordered-map-queue/queue"
	"log"
)

func main() {
	srv := queue.NewServer()
	if err := srv.Receive(); err != nil {
		log.Fatal(err)
	}
}
