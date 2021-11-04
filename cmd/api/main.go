package main

import (
	"log"

	"github.com/darianfd99/httpApiProject/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
