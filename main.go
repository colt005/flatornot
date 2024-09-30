package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/colt005/flatornot/server"
)

func main() {

	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	s.RegisterRoutes()

	go func() {
		s.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	//TODO: Flush backlog to DB

	// Graceful shutdown
	s.ShutDown()

	fmt.Println("Server exiting gracefully")

}
