package main

import (
	"log"
	"github.com/rohzo/distributed-file-system/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":4000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}