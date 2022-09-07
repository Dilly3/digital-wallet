package main

import (
	"github.com/dilly3/wallet-api/controller"
	"log"
)

func main() {
	err := controller.Start()
	if err != nil {
		log.Fatal(err)
	}
}
