package main

import (
	"fmt"
	"log"
	"math/big"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/web3"
)

func main() {
	// Start a new Geth node in the background
	cmd := exec.Command("geth", "--datadir=./mydata", "--networkid=12345", "--rpc", "--rpcport=8545", "--rpcaddr=0.0.0.0")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geth node started")
	
	// Give time for the node to start
	time.Sleep(5 * time.Second)
	
	// Connect to the Geth node using web3
	web3, err := web3.NewWeb3(web3.NewWeb3Config().
		SetHTTPEndpoint("http://localhost:8545").
		SetMaxRetries(3).
		SetRetryInterval(5))
	if err != nil {
		log.Fatal(err)
	}

	//configure the miner
	web3.Eth.SetMining(true)
	web3.Eth.SetGasPrice(big.NewInt(40000000000))

	// Monitor the status of the node
	blockNumber, err := web3.Eth.BlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block number:", blockNumber)

}
