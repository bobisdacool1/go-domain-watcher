package main

import (
	"go-domain-watcher/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
