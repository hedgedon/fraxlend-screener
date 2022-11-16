package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"log"
	"os"
	"os/signal"
)

type Scope struct {
	ctx       context.Context
	sdk       *thirdweb.ThirdwebSDK
	contracts map[string]*thirdweb.SmartContract
	tokenList map[string]string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx, cancel := context.WithCancel(context.Background())
	rpcUrl := os.Getenv("RPC_URL")
	sdk, err := thirdweb.NewThirdwebSDK(rpcUrl, nil)
	if err != nil {
		panic(err)
	}

	pretty(TokenPairList)

	tokenContract, err := sdk.GetToken(CHZ_TOKEN)
	token, err := tokenContract.BalanceOf(os.Getenv("MY_ADDRESS"))
	if err != nil {
		panic(err)
	}
	balance := token.DisplayValue
	fmt.Println("balance:", balance)

	scope := &Scope{
		sdk:       sdk,
		ctx:       ctx,
		tokenList: TokenPairList,
	}

	fraxlendPairContracts := scope.createPairContracts(FraxlendPairList)
	//pretty2(fraxlendPairContracts)
	b, err := json.MarshalIndent(fraxlendPairContracts, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
	scope.contracts = fraxlendPairContracts

	go scope.run()
	//go scope.run_cli()

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
	//program := cli.Run_cli()
	//program.Start()

	for i, _ := range s.contracts {
		s.getPairData(i)
	}

	s.getUserData(FXS_FRAX_POOL, os.Getenv("MY_ADDRESS"))
}

//func (s *Scope) run_cli() {
//	program := cli.Run_cli()
//	program.Start()
//}
