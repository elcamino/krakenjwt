package main

import (
	"log"

	kj "github.com/elcamino/krakenjwt"
)

func main() {
	k, err := kj.New("127.0.0.1:4680")
	if err != nil {
		log.Fatal(err)
	}

	k.Run()
}
