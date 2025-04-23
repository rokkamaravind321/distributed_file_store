package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents a peer(remote node) connection over TCP
type TCPPeer struct {
	//underlying connection
	conn net.Conn

	//whether the peer is outbound or inbound(dial&retrieve or accept&retrieve)
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportConfig struct {
	ListenAddress string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportConfig
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(config TCPTransportConfig) *TCPTransport {
	return &TCPTransport{
		TCPTransportConfig: config,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("TCP Transport: Accept error:", err)
		}

		fmt.Println("Incoming Connection :", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Println("TCP Transport: Handshake error:", err)
		return
	}

	// read loop
	msg := &Temp{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("TCP Transport: Decode error:", err)
			continue
		}
	}
}
