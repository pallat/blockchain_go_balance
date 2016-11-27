//go:generate abigen --sol vinTrack.sol --pkg main --out token.go

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"

    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/crypto"
)

func main() {
	c,err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ec := ethclient.NewClient(c)
	addr := common.HexToAddress("0x4df9bca2535a07cada6f73e5baa8f01b88655bea")
	b,err := ec.BalanceAt(ctx,  addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
	newContract()
}

func newContract() {
   	key, _ := crypto.GenerateKey()
    auth := bind.NewKeyedTransactor(key)

    sim := backends.NewSimulatedBackend(core.GenesisAccount{Address: auth.From, Balance: big.NewInt(10000000000)})

    // Deploy a token contract on the simulated blockchain
    _, _, token, err := DeployVinTracker(auth, sim)//, new(big.Int), "Simulated blockchain tokens", 0, "SBT")
    if err != nil {
        log.Fatalf("Failed to deploy new token contract: %v", err)
    }

	token.AddEvent(auth,"12345678","new")
	// t,_ := token.GetEventsByVIN(auth,"12345678")
	// fmt.Println(t)
	id,_ := token.AddEvent(auth,"12345678","sold")
	t,_ := token.GetEventById(auth,id.ChainId())
	fmt.Println(t)
	token.AddEvent(auth,"12345678","new owner")
	token.AddEvent(auth,"12345678",".....")
	t,_ = token.GetEventsByVIN(auth,"12345678")
	fmt.Println(t)

    // Print the current (non existent) and pending name of the contract
    // name, _ := token.Name(nil)
    // fmt.Println("Pre-mining name:", name)

    // name, _ = token.Name(&bind.CallOpts{Pending: true})
    // fmt.Println("Pre-mining pending name:", name)

    // // Commit all pending transactions in the simulator and print the names again
    // sim.Commit()

    // name, _ = token.Name(nil)
    // fmt.Println("Post-mining name:", name)

    // name, _ = token.Name(&bind.CallOpts{Pending: true})
    // fmt.Println("Post-mining pending name:", name)

}