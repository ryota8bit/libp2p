package service

import (
	"bufio"
	"context"
	"fmt"
	"log"

	host "github.com/libp2p/go-libp2p-host"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-protocol"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	P2PProtocolID protocol.ID = "/p2p/1.0.0"
)

func StreamHandler(ha host.Host) {
	log.Println("listening for connections")
	ha.SetStreamHandler(P2PProtocolID, HandleStream)
	select {}
}

func StreamHandlerWithTarget(ha host.Host, target string) {
	// Stream handler that opens when specified as a target
	ha.SetStreamHandler(P2PProtocolID, HandleStream)

	// Open a new stream to given peer p, and writes a p2p/protocol
	stream := newStream(ha, target)

	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(
		bufio.NewReader(stream),
		bufio.NewWriter(stream),
	)

	// Create a my thread to stream.
	go WriteBlockChain(rw)
	go WriteBlockChainWithInputPrompt(rw)
	go ReadBlockChain(rw)

	select {} // hang forever
}

// newStream open a new stream to given peer p, and writes a p2p/protocol
func newStream(ha host.Host, target string) inet.Stream {
	peerID, targetAddr := getPeerIDWithMultiaddr(target)
	ha.Peerstore().AddAddr(
		peerID,
		targetAddr,
		pstore.PermanentAddrTTL,
	)
	log.Println("opening stream")
	log.Println("We are ", ha.Peerstore().Peers())

	// make a new stream from host B to host A
	s, err := ha.NewStream(context.Background(), peerID, P2PProtocolID)
	if err != nil {
		log.Fatalln(err)
	}
	return s
}

func getPeerIDWithMultiaddr(target string) (peer.ID, ma.Multiaddr) {
	ipfsAddr, err := ma.NewMultiaddr(target)
	if err != nil {
		log.Fatalln(err)
	}

	pid, err := ipfsAddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Fatalln(err)
	}

	peerID, err := peer.IDB58Decode(pid)
	if err != nil {
		log.Fatalln(err)
	}

	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerID)))
	return peerID, ipfsAddr.Decapsulate(targetPeerAddr)
}
