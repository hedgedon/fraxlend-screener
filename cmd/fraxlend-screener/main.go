package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	RPC_URL := os.Getenv("RPC_URL")
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")

	ctx := context.Background()

	sdk, err := thirdweb.NewThirdwebSDK(RPC_URL, &thirdweb.SDKOptions{
		PrivateKey: PRIVATE_KEY,
	})
	if err != nil {
		panic(err)
	}

	contract, err := sdk.GetContractFromAbi(CRV_FRAX_POOL, FRAXLEND_PAIR_ABI)
	if err != nil {
		panic(err)
	}
	fmt.Println(contract)

	name, err := contract.Call(ctx, "name")
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", name)

	symbol, err := contract.Call(ctx, "symbol")
	if err != nil {
		panic(err)
	}
	fmt.Println("symbol:", symbol)

	totalBorrow, err := contract.Call(ctx, "totalBorrow")
	if err != nil {
		panic(err)
	}
	fmt.Println("total borrow: ", totalBorrow)
	// total borrow:  [3521540165111577988854598 3508827027810177699518635]
	//TOTAL BORROW VALUE $3.52m

	totalCollateral, err := contract.Call(ctx, "totalCollateral")
	if err != nil {
		panic(err)
	}
	fmt.Println("total Collateral: ", totalCollateral)

	totalAsset, err := contract.Call(ctx, "totalAsset")
	if err != nil {
		panic(err)
	}
	fmt.Println("total Asset: ", totalAsset)

	totalSupply, err := contract.Call(ctx, "totalSupply")
	if err != nil {
		panic(err)
	}
	fmt.Println("total Supply: ", totalSupply)

}
