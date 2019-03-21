package service

import (
	"bufio"
	"log"

	net "github.com/libp2p/go-libp2p-net"
)

// HandleStream. stream to open when specified in target
// stream 's' will stay open until you close it (or the other side closes it).
func HandleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go ReadBlockChain(rw)
	go WriteBlockChain(rw)
	go WriteBlockChainWithInputPrompt(rw)
}
