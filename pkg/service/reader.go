package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// ReadBlockChain read incoming blockchain from a peer and stdout.
func ReadBlockChain(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}

		if str != "\n" {

			chain := make([]Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			if len(chain) > len(BlockChain) {
				BlockChain = chain
				validateBlockChain()

				bytes, err := json.MarshalIndent(BlockChain, "", "  ")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()
		}
	}
}

func validateBlockChain() {
	for idx, v := range BlockChain {
		if len(BlockChain) > idx+1 && !IsBlockValid(BlockChain[idx+1], v) {
			log.Fatal(errors.New("invalid block"))
		}
	}
}
