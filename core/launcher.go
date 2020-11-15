package core

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	golog "github.com/ipfs/go-log"
	peer "github.com/libp2p/go-libp2p-core/peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
	gologging "github.com/whyrusleeping/go-logging"
)

//Launch starts the ledger
func (b *Chain) Launch() (*Chain, error) {

	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), "genesis block", genesisBlock.CalculateHash(), ""}

	b.BlockChain = append(b.BlockChain, genesisBlock)

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	// Parse options
	address := &b.P2pConfig.Address
	port := &b.P2pConfig.Port
	target := &b.P2pConfig.Target
	secio := &b.P2pConfig.Secio
	seed := &b.P2pConfig.Seed

	flag.Parse()

	if *port == 0 {
		log.Fatal("Please provide a port to bind on with -l")

	}

	// Make a host that listens on the given multiaddress
	ha, err := MakeBasicHost(*address, *port, *secio, *seed)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if *target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is
		// a user-defined protocol name.
		ha.SetStreamHandler("/p2p/1.0.0", b.HandleStream)

		select {} // hang forever
		/**** This is where the listener code ends ****/
	} else {
		ha.SetStreamHandler("/p2p/1.0.0", b.HandleStream)

		// The following code extracts target's peer ID from the
		// given multiaddress
		ipfsaddr, err := ma.NewMultiaddr(*target)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		peerid, err := peer.IDB58Decode(pid)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		// Decapsulate the /ipfs/<peerID> part from the target
		// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
		targetPeerAddr, _ := ma.NewMultiaddr(
			fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
		targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

		// We have a peer ID and a targetAddr so we add it to the peerstore
		// so LibP2P knows how to contact it
		ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

		log.Println("opening stream")
		// make a new stream from host B to host A
		// it should be handled on host A by the handler we set above because
		// we use the same /p2p/1.0.0 protocol
		s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		// Create a buffered stream so that read and writes are non blocking.
		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		//go b.Run(":8090")

		// Create a thread to read and write data.
		go b.WriteDataCli(rw)
		go b.ReadDataCli(rw)

		select {} // hang forever

		return b, nil
	}
}
