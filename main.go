package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"snake-tail/config"
	"snake-tail/http/server"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	httpServer *server.Server
)

func main() {

	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	httpServer = server.NewServer(config.App)

	// Start http server
	err = httpServer.Start()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C), SIGKILL, SIGQUIT or SIGTERM (Ctrl+/)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	// Block until we receive our signal
	signal := <-c
	fmt.Println("bootstrap.init.Start", fmt.Sprintf("Received Signal: %s", signal))

	// Start destructing the process
	httpServer.Stop()
}
