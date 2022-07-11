package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"github.com/rrobrms/w3w/pkg/client"
	"github.com/rrobrms/w3w/pkg/utils"
	"github.com/spf13/cobra"
)

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Get transaction hash info",
	Long:  `This command fetches and decode the transaction info from the given transaction hash.`,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			chainId, _ = cmd.Flags().GetString("chainid")
			txHash     = args[0]
		)

		if len(args) == 0 {
			fmt.Println("Please provide a transaction hash")
			return
		} else if len(args) > 1 {
			fmt.Println("Please provide only one transaction hash")
			return
		}

		if chainId != "" {
			cl, err := client.GetClient(chainId)
			if err != nil {
				log.Fatal("failed to connect to ethereum client", err)
			}
			cle, err := client.GetClientEthScan(chainId)
			if err != nil {
				log.Fatal("failed to connect to etherScan client", err)
			}
			getTxHash(cl, cle, txHash, chainId)
		} else {
			cl, err := client.GetClient(chainId)
			if err != nil {
				log.Fatal("failed to connect to ethereum client", err)
			}
			cle, err := client.GetClientEthScan(chainId)
			if err != nil {
				log.Fatal("failed to connect to etherScan client", err)
			}
			getTxHash(cl, cle, txHash, chainId)
		}
	},
}

func init() {
	rootCmd.AddCommand(txCmd)

	txCmd.PersistentFlags().String("chainid", "", "Get tx info from a specific chain")
}

func ParseTransactionInfos(tx *types.Transaction, msg types.Message, receipt *types.Receipt, isContract bool) {

	var (
		from     = msg.From().Hex()
		to       = tx.To().Hex()
		value    = tx.Value().String()
		cost     = utils.WeiToEther(tx.Cost())
		GasPrice = utils.WeiToEther(tx.GasPrice())
		GasLimit = tx.Gas()
		GasUsed  = receipt.GasUsed
		dataHex  = hex.EncodeToString(tx.Data())
	)

	fmt.Printf("BlockNumber: %v | Chain ID: %v\n", receipt.BlockNumber, tx.ChainId())
	fmt.Printf("From: %s\n", from)
	fmt.Printf("To: %s (is a contract: %t)\n", to, isContract)
	fmt.Printf("Value: %s\n", value)
	fmt.Printf("Transaction Fee: %v Ether\n", cost)
	fmt.Printf("Gas Price: %v Gwei\n", GasPrice)
	fmt.Printf("Gas Limit/Gas Used: %v | %v\n", GasLimit, GasUsed)
	fmt.Printf("Type: %v | Nonce: %v | tx index: %v\n", receipt.Type, msg.Nonce(), receipt.TransactionIndex)
	fmt.Printf("Transaction Data in hex: %s\n", dataHex)
}

func getTxHash(client *ethclient.Client, clientEthScan *etherscan.Client, txHash string, chainId string) {

	_txHash := common.HexToHash(txHash)

	tx, isPending, err := client.TransactionByHash(context.Background(), _txHash)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), _txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx isPending: %t\n", isPending)
	fmt.Printf("Status: %d\n", receipt.Status)

	address := common.HexToAddress(tx.To().String())
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0

	source, err := clientEthScan.ContractSource(tx.To().String())
	if err != nil {
		log.Fatal(err)
	}
	proxy := source[0].Proxy

	if proxy != "0" {

		sourceImpl, err := clientEthScan.ContractSource(source[0].Implementation)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Contract Implementation of: %s\n", sourceImpl[0].ContractName)
		ParseTransactionInfos(tx, msg, receipt, isContract)

		contractABI, err := abi.JSON(strings.NewReader(sourceImpl[0].ABI))
		if err != nil {
			log.Fatal(err)
		}
		utils.DecodeTransactionInputData(contractABI, tx.Data())
		utils.DecodeTransactionLogs(receipt, contractABI)

	} else {
		fmt.Printf("Contract: %s\n", source[0].ContractName)
		ParseTransactionInfos(tx, msg, receipt, isContract)

		contractABI, err := abi.JSON(strings.NewReader(source[0].ABI))
		if err != nil {
			log.Fatal(err)
		}
		utils.DecodeTransactionInputData(contractABI, tx.Data())
		utils.DecodeTransactionLogs(receipt, contractABI)
	}

}
