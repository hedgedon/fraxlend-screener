package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
	"log"
	"os"
	"os/signal"
)

type Scope struct {
	ctx       context.Context
	sdk       *thirdweb.ThirdwebSDK
	contracts map[string]*thirdweb.SmartContract
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx, cancel := context.WithCancel(context.Background())
	RpcUrl := os.Getenv("RPC_URL")
	//PrivateKey := os.Getenv("PRIVATE_KEY")

	sdk, err := thirdweb.NewThirdwebSDK(RpcUrl, nil)
	if err != nil {
		panic(err)
	}

	scope := &Scope{
		sdk: sdk,
		ctx: ctx,
	}
	fraxlendPairContracts := scope.createPairContracts()
	scope.contracts = fraxlendPairContracts

	go scope.run()

	// if u want to run go scope.run() as a goroutine, u need the part below.
	// otherwise, call with scope.run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		signal.Stop(c)
		cancel()
		<-ctx.Done()
		log.Println("~ ~ ~ ~ exiting ~ ~ ~ ~")
		os.Exit(0)
	}
}

func (s *Scope) run() {
	s.getPairData()
	s.getUserData()
}
