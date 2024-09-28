package main

import (
	"log"

	"github.com/colt005/flatornot/server"
)

func main() {

	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	s.RegisterRoutes()

	s.Start()

}
