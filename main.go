package main

import (
	"fmt"
	"log"

	"github.com/rokkamaravind321/distributed_file_store/p2p"
)

func onPeer(peer p2p.Peer) error {
	fmt.Println("onPeer", peer)
	return nil
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        onPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("RPC: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

}
