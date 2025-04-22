package main

import (
	"fmt"
	"log"

	"github.com/rokkamaravind321/distributed_file_store/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

	fmt.Println("Hello, World!")
}
