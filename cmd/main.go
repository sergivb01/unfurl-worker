package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/sergivb01/unfurl-worker/internal/encoder"
	"github.com/sergivb01/unfurl-worker/internal/meta"
	"github.com/sergivb01/unfurl-worker/internal/metaclient"
)

type Request struct {
	URL  string
	Sent time.Time
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect to nats: %s", err)
	}
	nats.RegisterEncoder("msgpack", &encoder.MsgPackEncoder{})

	c, err := nats.NewEncodedConn(nc, "msgpack")
	if err != nil {
		log.Fatalf("failed to create encoded connection: %s", err)
	}

	go func() {
		sendMessage(c)
	}()

	c.QueueSubscribe("requests", "data", func(subj, reply string, r *Request) {
		fmt.Printf("Received %+v\n", r)

		reader, err := metaclient.GetReaderFromURL(context.TODO(), r.URL, true)
		if err != nil {
			panic(err)
		}

		info, _ := meta.ExtractInfoFromReader(reader)

		fmt.Printf("responding with:\n%s\n", info)

		if err := c.Publish(reply, info); err != nil {
			log.Printf("failed to reply: %s", err)
		}
	})

	if err := nc.Flush(); err != nil {
		log.Printf("err trying to flush: %s", err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", "requests")
	log.SetFlags(log.LstdFlags)

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println()

	log.Printf("Draining...")
	if err := nc.Drain(); err != nil {
		log.Printf("error trying to drain: %s", err)
	}

	log.Fatalf("Exiting")
}

func sendMessage(nc *nats.EncodedConn) {
	fmt.Println("sleeping 3s")
	time.Sleep(time.Second * 3)
	fmt.Println("sending")

	start := time.Now()

	var info meta.PageInfo
	if err := nc.Request("requests", &Request{URL: "https://sergivos.dev"}, &info, time.Second*3); err != nil {
		log.Printf("err publishing: %s", err)
		return
	}
	fmt.Printf("response: %v\n", info)

	fmt.Printf("took %s to receive message", time.Since(start))
}
