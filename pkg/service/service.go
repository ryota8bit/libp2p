package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"sync"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
	host "github.com/libp2p/go-libp2p-host"
	ma "github.com/multiformats/go-multiaddr"
)

var mutex = &sync.Mutex{}

// InitBlockChain Block initialization.
func InitBlockChain() {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, CalculateHash(genesisBlock), ""}
	BlockChain = append(BlockChain, genesisBlock)
}

// BasicHost creates a LibP2P host with a random peer ID listening on the
func BasicHost(listenPort int, randseed int64) (host.Host, error) {

	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}

	p2pHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	// Build host multiaddress
	multiAddr := fmt.Sprintf("/ipfs/%s", p2pHost.ID().Pretty())
	hostAddr, _ := ma.NewMultiaddr(multiAddr)

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := p2pHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
	log.Printf("Now run \"p2p-app -l %d -d /ip4/127.0.0.1/tcp/%d%s\" on a different terminal\n", listenPort+1, listenPort, multiAddr)
	return p2pHost, nil
}
