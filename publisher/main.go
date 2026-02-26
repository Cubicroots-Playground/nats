package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	for i := range 100 {
		fmt.Println("publish message")

		// Publish simple message.
		nc.Publish("my_subject1", []byte(fmt.Sprintf("Hello World %d", i)))

		// Send a request.
		m, err := nc.Request("my_subject2", []byte(fmt.Sprintf("Hello World %d", i)), time.Second)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Received a message: %s\n", string(m.Data))
		}

		time.Sleep(time.Second)
	}

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
