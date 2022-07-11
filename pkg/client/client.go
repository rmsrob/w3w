package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
)

func GetClient(c string) (client *ethclient.Client, err error) {

	k := ""
	var url string

	if c == "" {
		url = `https://` + "mainnet" + `.infura.io/v3/` + k
	} else if c == "rinkeby" {
		url = `https://` + "rinkeby" + `.infura.io/v3/` + k
	} else if c == "goerli" {
		url = `https://` + "goerli" + `.infura.io/v3/` + k
	} else if c == "local" {
		url = `http://` + "localhost/" + k
	}

	client, err = ethclient.Dial(url)
	if err != nil {
		log.Fatal("failed to connect to ethereum client", err)
		return
	}
	defer client.Close()

	return client, err
}

func GetClientEthScan(c string) (client *etherscan.Client, err error) {
	etherscanAPIKEY := ""
	client = etherscan.New(etherscan.Rinkby, etherscanAPIKEY)

	return client, err
}
