package main

import (
	"groupie/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

}
