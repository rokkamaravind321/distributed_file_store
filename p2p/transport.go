package p2p

// Interface that represent remote node(TCP,UDP,Websockets)
type Peer interface {
	Close() error
}

// Handle communication between the nodes and the network of any from (TCP,UDP,Websockets)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
