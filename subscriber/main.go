package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	sub, err := nc.QueueSubscribe("my_subject1", "my_queue", func(m *nats.Msg) {
		fmt.Printf("Received a message (worker 1): %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}
	sub2, err := nc.QueueSubscribe("my_subject1", "my_queue", func(m *nats.Msg) {
		fmt.Printf("Received a message (worker 2): %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}

	// Simple Async Subscriber with reply.
	sub3, err := nc.Subscribe("my_subject2", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		nc.Publish(m.Reply, []byte(fmt.Sprintf("Reply to '%s'", string(m.Data))))
	})
	if err != nil {
		panic(err)
	}

	// Async sub is running in another routine, just wait here.
	time.Sleep(time.Second * 10)

	// Drain subscriber.
	err = sub.Drain()
	if err != nil {
		panic(err)
	}
	err = sub2.Drain()
	if err != nil {
		panic(err)
	}
	err = sub3.Drain()
	if err != nil {
		panic(err)
	}

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
