package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/KoganezawaRyouta/libp2p/pkg/service"
	"github.com/libp2p/go-libp2p-host"
	"github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Block initialization
	service.InitBlockChain()

	// Make a host that listens on the given multiaddress
	ha, err := service.MakeBasicHost(*listenF, *seed)
	if err != nil {
		log.Fatal(err)
	}

	// set stream handler
	if *target == "" {
		setStreamHandler(ha)
	} else {
		setStreamHandlerWithTarget(ha, *target)
	}
}

func setStreamHandler(ha host.Host) {
	log.Println("listening for connections")

	// Set a stream handler on host A. /p2p/1.0.0 is
	// a user-defined protocol name.
	ha.SetStreamHandler("/p2p/1.0.0", service.HandleStream)

	select {} // hang forever
	/**** This is where the listener code ends ****/
}

func setStreamHandlerWithTarget(ha host.Host, target string) {
	ha.SetStreamHandler("/p2p/1.0.0", service.HandleStream)

	// The following code extracts target's peer ID from the
	// given multiaddress
	ipfsaddr, err := ma.NewMultiaddr(target)
	if err != nil {
		log.Fatalln(err)
	}

	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Fatalln(err)
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		log.Fatalln(err)
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")
	log.Println("We are ", ha.Peerstore().Peers())

	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /p2p/1.0.0 protocol
	s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
	if err != nil {
		log.Fatalln(err)
	}
	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Create a my thread to read and write data.
	go service.WriteBlockChain(rw)
	go service.WriteBlockChainWithInputPrompt(rw)
	go service.ReadBlockChain(rw)

	select {} // hang forever
}
