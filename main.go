package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
)

func main() {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	var isValid string
	if re.MatchString(address) {
		isValid = color.GreenString("valid")
	} else {
		isValid = color.RedString("invalid")
	}

	t := color.New(color.FgCyan,color.Bold)
	g := color.New(color.FgHiWhite)

	var (
		verbose bool = false
		verbose_usage string = "Get nice output"
	)
	flag.BoolVar(&verbose, "v", verbose, verbose_usage)
    flag.Parse()

	if verbose {
		t.Println("╭─────────────┬──────────────────────────────────────────────────────────────────╮")
		t.Printf("│ private key │ %v │\n", privateKeyHex)
		t.Println("├─────────────├──────────────────────────────────────────────────────────────────┤")
		t.Printf("│ address     │ %v                       │\n", address)
		t.Println("╰─────────────┴──────────────────────────────────────────────────────────────────╯")
		g.Printf("  %s   https://etherscan.io/address/%v \n", isValid, address)	
	} else {
		fmt.Println(privateKeyHex, address)
	}
}
