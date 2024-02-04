package main

import (
	"github.com/therealdannzor/go-ordered-map-queue/queue"
	"log"
)

func main() {
	err := queue.Send()
	if err != nil {
		log.Fatal(err)
	}
}
