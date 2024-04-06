package main

import (
	"fmt"
	"os"
	"os/signal"
	"snake-tail/config"
	"snake-tail/http/server"
	"snake-tail/workers"
	"syscall"
)

var (
	httpServer *server.Server
)

func main() {

	workers.InitWorkers()

	httpServer = server.NewServer(config.App)

	// Start http server
	err := httpServer.Start()
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
