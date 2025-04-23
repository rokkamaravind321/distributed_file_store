package main

import (
	"log"

	"github.com/rokkamaravind321/distributed_file_store/p2p"
)

func main() {

	config := p2p.TCPTransportConfig{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GobDecoder{},
	}

	tr := p2p.NewTCPTransport(config)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

}
