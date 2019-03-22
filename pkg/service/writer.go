package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// WriteBlockChain propagate the acquired block data to another node.
func WriteBlockChain(rw *bufio.ReadWriter) {
	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			mutex.Lock()
			bytes, err := json.Marshal(BlockChain)
			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()
			mutex.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mutex.Unlock()

		}
	}()
}

// WriteBlockChainWithInputPrompt Get block from input prompt and propagate block to another node.
func WriteBlockChainWithInputPrompt(rw *bufio.ReadWriter) {

	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		point := getPoint(sendData)
		newBlock := GenerateBlock(BlockChain[len(BlockChain)-1], point)
		bytes := appendBlockWithMarshal(newBlock)
		spew.Dump(BlockChain)

		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()
	}
}

func appendBlockWithMarshal(newBlock Block) []byte {
	if IsBlockValid(newBlock, BlockChain[len(BlockChain)-1]) {
		mutex.Lock()
		BlockChain = append(BlockChain, newBlock)
		mutex.Unlock()
	}
	return marshalWithDumpBlockChain()
}

func marshalWithDumpBlockChain() []byte {
	bytes, err := json.Marshal(BlockChain)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func getPoint(sendData string) int {
	sendData = strings.Replace(sendData, "\n", "", -1)
	point, err := strconv.Atoi(sendData)
	if err != nil {
		log.Fatal(err)
	}
	return point
}
