package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	net "github.com/libp2p/go-libp2p-core/network"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

//HandleStream handle stream
func (b *Chain) HandleStream(s net.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go b.ReadDataCli(rw)
	go b.WriteDataCli(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}

// ReadData read data
func (b *Chain) ReadDataCli(rw *bufio.ReadWriter) {
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
			if len(chain) > len(b.BlockChain) {
				b.BlockChain = chain
				bytes, err := json.MarshalIndent(b.BlockChain, "", "  ")
				if err != nil {

					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()
		}
	}
}

// WriteData write data
func (b *Chain) WriteDataCli(rw *bufio.ReadWriter) {

	go func() {
		for {

			//displaying blockchain content to handle 3 client

			time.Sleep(5 * time.Second)
			mutex.Lock()
			bytes, err := json.Marshal(b.BlockChain)
			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()

			mutex.Lock()
			_, err = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))

			if err != nil {
				log.Println(err)
			}
			err = rw.Flush()

			if err != nil {
				log.Println(err)
			}
			mutex.Unlock()

		}
	}()

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sentData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sentData = strings.Replace(sentData, "\n", "", -1)

		oldBlock := b.BlockChain[len(b.BlockChain)-1]
		newBlock := oldBlock.GenerateBlock(sentData)

		if newBlock.IsBlockValid(b.BlockChain[len(b.BlockChain)-1]) {
			mutex.Lock()
			b.BlockChain = append(b.BlockChain, newBlock)
			mutex.Unlock()
		}
		bytes, err := json.Marshal(b.BlockChain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(b.BlockChain)

		mutex.Lock()
		_, err = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		if err != nil {
			log.Println(err)
		}

		err = rw.Flush()
		if err != nil {
			log.Println(err)
		}

		mutex.Unlock()
	}

}

func (b *Chain) WriteDataRest(c *gin.Context) {

	fmt.Print("> ")
	sentData := c.Query("data")
	if sentData == "" {
		log.Fatal("you sent an empty string")
		return
	}

	sentData = strings.Replace(sentData, "\n", "", -1)
	oldBlock := b.BlockChain[len(b.BlockChain)-1]
	newBlock := oldBlock.GenerateBlock(sentData)

	if newBlock.IsBlockValid(b.BlockChain[len(b.BlockChain)-1]) {
		mutex.Lock()
		b.BlockChain = append(b.BlockChain, newBlock)
		mutex.Unlock()

	}
}

func (b *Chain) ReadDataRest(c *gin.Context) {
	c.JSON(200, b.BlockChain)

}
