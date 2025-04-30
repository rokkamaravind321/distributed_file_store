package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {

	opts := TCPTransportOpts{
		ListenAddress: ":4000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}

	tr := NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddress, opts.ListenAddress)

	assert.Nil(t, tr.ListenAndAccept())
	
}
