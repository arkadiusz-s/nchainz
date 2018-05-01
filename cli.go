package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// For processing cmd line arguments
type CLI struct {
}

//
// Main method to parse and process cmds
//
func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		cli.printHelp()
		os.Exit(1)
	}

	// Use flag package to parse cmd line arguments
	walletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	printAddressesCmd := flag.NewFlagSet("printaddresses", flag.ExitOnError)
	nodeCmd := flag.NewFlagSet("node", flag.ExitOnError)
	createBCCmd := flag.NewFlagSet("createbc", flag.ExitOnError)

	bcAddress := createBCCmd.String("address", "", "Genesis reward sent to this address.")

	switch os.Args[1] {
	case "createwallet":
		err := walletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printaddresses":
		err := printAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "node":
		err := nodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createbc":
		err := createBCCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printHelp()
		os.Exit(1)
	}

	if walletCmd.Parsed() {
		cli.createWallet()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if nodeCmd.Parsed() {
		cli.node(os.Args)
	}

	if printAddressesCmd.Parsed() {
		cli.printAddresses()
	}

	if createBCCmd.Parsed() {
		if *bcAddress == "" {
			createBCCmd.Usage()
			os.Exit(1)
		}
		cli.createBC(*bcAddress)
	}
}

func (cli *CLI) printHelp() {
	fmt.Println("===== Help menu =====")
	fmt.Println("go run *.go createwallet                  --- Creates a wallet with a pair of keys")
	fmt.Println("go run *.go createbc -address ADDRESS     --- Creates new blockchain. ADDRESS gets genesis reward")
	fmt.Println("go run *.go printchain                    --- Print all the blocks in the blockchain")
	fmt.Println("go run *.go printaddresses								 --- Lists all addresses in walletFile")
	fmt.Println("go run *.go node                          --- start up a full node")
}

func (cli *CLI) getBalance(address string) {
	/*bc := NewBlockchain(address)
	defer bc.db.Close()*/
}

func (cli *CLI) printChain() {
	bcs := CreateNewBlockchains("blockchain.db")
	bc := bcs.GetChain(MATCH_CHAIN)
	defer bc.db.Close()

	bci := bc.Iterator()
	fmt.Println(MATCH_CHAIN)
	fmt.Println("Height:", bc.height)
	for {
		block, _ := bci.Next()

		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Validated Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("-------------------------------")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	bc = bcs.GetChain(NATIVE_CHAIN)

	bci = bc.Iterator()
	fmt.Println(NATIVE_CHAIN)
	fmt.Println("Height:", bc.height)
	for {
		block, _ := bci.Next()

		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Validated Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("-------------------------------")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) printAddresses() {
	ws := NewWalletStore()
	addresses := ws.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func (cli *CLI) createWallet() {
	ws := NewWalletStore()
	address := ws.AddWallet()
	ws.Persist()

	fmt.Printf("New wallet's address: %s\n", address)
}

func (cli *CLI) node(args []string) {
	// TODO: parse args properly
	port, _ := strconv.Atoi(args[2])
	seed := args[3]

	StartNode(uint(port), seed)
}

func (cli *CLI) createBC(address string) {
	bcs := CreateNewBlockchains("blockchain.db")
	bc := bcs.GetChain(NATIVE_CHAIN)
	transfer := Transfer{
		Amount:      10,
		FromAddress: "Satoshi",
		ToAddress:   "Negansoft",
		Signature:   nil,
	}
	tokenData := TokenData{
		Orders:     nil,
		ClaimFunds: nil,
		Transfers:  []Transfer{transfer},
	}
	block := NewBlock(tokenData, TOKEN_BLOCK, bc.GetTipHash())
	bcs.AddBlock(NATIVE_CHAIN, *block)
}
