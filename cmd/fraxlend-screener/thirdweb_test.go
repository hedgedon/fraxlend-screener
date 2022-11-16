package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"log"
	"os"
	"testing"
)

func TestContractEvent(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Errorf("error: %v \n", err)
	}
	contract, err := sdk.GetContractFromAbi(FRAX_TOKEN, FRAX_CONTRACT_ABI)
	filters := map[string]interface{}{
		"from": common.HexToAddress(os.Getenv("MY_ADDRESS")),
	}
	queryOptions := thirdweb.EventQueryOptions{
		Filters: filters,
	}
	// Now we can query for the Transfer events
	events, _ := contract.Events.GetEvents(context.Background(), "Transfer", queryOptions)
	if err != nil {
		t.Fatal(err)
	}
	for i, e := range events {
		t.Logf("%v) Event: %v. %v", i, e.EventName, e.Transaction)
	}
}
