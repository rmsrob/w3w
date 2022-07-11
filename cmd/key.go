package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Get private key tools",
	Long:  `This command provides tools for getting private keys.`,
	Run: func(cmd *cobra.Command, args []string) {

		priv, _ := cmd.Flags().GetString("priv")
		// v, _ := cmd.Flags().GetString("v")

		if priv != "" {
			adr := deriveAddress(priv)
			fmt.Printf(color.GreenString("Derivated address: %s\n"), adr)
		} else {
			privateKeyHex, address := createKeyPair()
			// fmt.Println(privateKeyHex, address)
			t := color.New(color.FgCyan, color.Bold)
			g := color.New(color.FgHiWhite)
			t.Println("╭─────────────┬──────────────────────────────────────────────────────────────────╮")
			t.Printf("│ private key │ %v │\n", privateKeyHex)
			t.Println("├─────────────├──────────────────────────────────────────────────────────────────┤")
			t.Printf("│ address     │ %v                       │\n", address)
			t.Println("╰─────────────┴──────────────────────────────────────────────────────────────────╯")
			g.Printf(" https://etherscan.io/address/%s\n", address)

		}
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	keyCmd.PersistentFlags().String("v", "", "Create key pair with color verbose.")
	keyCmd.PersistentFlags().String("priv", "", "Create key pair from a private key, or derive address from a it.")
}

func createKeyPair() (privateKey string, address string) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	adr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyHex, adr
}

/*
https://www.securityevaluators.com/casestudies/ethercombing/
The Ethereum project uses elliptic curve cryptography to generate the public/private key pair.
[256-bit private key] --> secp256k1 ECDSA curve --> [public key]
[public key] --> keccak256 & truncated to the lower 160 bits --> Ethereum address.
*/
func deriveAddress(priv string) (address string) {

	secp256k1N, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	// decode simulated hex string to bytes
	privateKeyBytes, err := hexutil.Decode("0x" + priv)
	if err != nil {
		log.Fatal(err)
	}

	privt := new(ecdsa.PrivateKey)
	privt.PublicKey.Curve = crypto.S256()
	if 8*len(privateKeyBytes) != privt.Params().BitSize {
		fmt.Printf("invalid length, need %d bits", privt.Params().BitSize)
	}
	privt.D = new(big.Int).SetBytes(privateKeyBytes)

	if privt.D.Cmp(secp256k1N) >= 0 {
		fmt.Println("invalid private key, >=N")
	}

	if privt.D.Sign() <= 0 {
		fmt.Println("invalid private key, zero or negative")
	}

	privt.PublicKey.X, privt.PublicKey.Y = privt.PublicKey.Curve.ScalarBaseMult(privateKeyBytes)
	if privt.PublicKey.X == nil {
		log.Fatal("invalid private key")
	}

	adr := crypto.PubkeyToAddress(privt.PublicKey).Hex()

	return adr
}
