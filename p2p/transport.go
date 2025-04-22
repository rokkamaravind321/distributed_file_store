package p2p

// Interface that represent remote node(TCP,UDP,Websockets)
type Peer interface {
}

// Handle comminucation between the nodes and the network
type Transport interface {
	ListenAndAccept() error
}
