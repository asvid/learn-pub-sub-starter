package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	connectionStr := "amqp://guest:guest@localhost:5672/"

	con, _ := amqp.Dial(connectionStr)
	defer func(con *amqp.Connection) {
		_ = con.Close()
	}(con)
	fmt.Println("Connected to RabbitMQ")

	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify the channel when receiving SIGINT (Ctrl+C) or SIGTERM
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-sigChan
	fmt.Printf("Received signal %v. Shutting down...\n", sig)

	// The deferred connection close will happen automatically
	fmt.Println("Connection closed. Goodbye!")
}
