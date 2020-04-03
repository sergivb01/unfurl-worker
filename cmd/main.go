package main

import (
	"fmt"
	"time"

	"github.com/sergivb01/unfurl-worker/internal/queue"
)

func main() {
	q, err := queue.NewQueue(false)
	if err != nil {
		panic(err)
	}

	go func(q *queue.NATSQueue) {
		sendMessage(q)
	}(q)

	if err := q.Subscribe(); err != nil {
		panic(err)
	}

	q.Start()
}

func sendMessage(q *queue.NATSQueue) {
	s := time.Now()
	res, err := q.Queue("http://www.mytwogeeks.com/rsvpVideo.php")
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	fmt.Printf("Took %s to receive %#+v!\n", time.Since(s), res)
}
