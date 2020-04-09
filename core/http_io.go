package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func (b *Chain) WriteDataRest(c *gin.Context) {

	//for {
	fmt.Print("> ")
	//sentData, err := stdReader.ReadString('\n')
	sentData := c.Query("data")
	if sentData == "" {
		log.Fatal("you sent an empty string")
		return
	}
	log.Println("werery", sentData)
	sentData = strings.Replace(sentData, "\n", "", -1)

	newBlock := GenerateBlock(b.BlockChain[len(b.BlockChain)-1], sentData)

	if IsBlockValid(newBlock, b.BlockChain[len(b.BlockChain)-1]) {
		mutex.Lock()
		b.BlockChain = append(b.BlockChain, newBlock)
		mutex.Unlock()
		//}
		/*bytes, err := json.Marshal(b.BlockChain)
		if err != nil {
			log.Println(err)
		}

		spew.Dump(b.BlockChain)

		mutex.Lock()
		//_,err = rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		if err != nil {
			log.Println(err)
		}

		//err = rw.Flush()
		if err != nil {
			log.Println(err)
		}

		mutex.Unlock()*/
	}

}
