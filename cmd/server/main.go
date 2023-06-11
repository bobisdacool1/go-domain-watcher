package main

import (
	"log"
	"url-accessibility-checker/internal/server"
)

func main() {
	s := server.NewServer()
	if err := s.Listen(); err != nil {
		log.Fatal(err)
	}
}
