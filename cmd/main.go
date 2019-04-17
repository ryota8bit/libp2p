package main

import (
	"flag"
	"log"

	"github.com/KoganezawaRyouta/libp2p/pkg/service"
	golog "github.com/ipfs/go-log"
	gologging "github.com/whyrusleeping/go-logging"
)

func main() {
	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// Block initialization
	service.InitBlockChain()
	golog.SetAllLoggers(gologging.INFO)

	// Make a host that listens on the given multiaddress
	ha, err := service.BasicHost(*listenF)
	if err != nil {
		log.Fatal(err)
	}

	// set stream handler
	if *target == "" {
		service.StreamHandler(ha)
	} else {
		service.StreamHandlerWithTarget(ha, *target)
	}
}
