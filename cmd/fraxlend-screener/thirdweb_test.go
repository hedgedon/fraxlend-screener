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
	// First we define a filter to only get Transfer events where the "from" address is "0x..."
	// Note that you can only add filters for indexed parameters on events
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

func TestFraxlendContractEvents(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Errorf("error: %v \n", err)
	}
	contract, err := sdk.GetContractFromAbi(FXS_FRAX_POOL, FRAXLEND_PAIR_ABI)

	// First we define a filter to only get AddCollateral events where the "_borrower" address is "0x..."
	// Note that you can only add filters for indexed parameters on events
	filters := map[string]interface{}{
		"_borrower": common.HexToAddress(os.Getenv("MY_ADDRESS")),
	}
	queryOptions := thirdweb.EventQueryOptions{
		Filters: filters,
	}
	// Now we can query for the Transfer events
	collateralEvents, _ := contract.Events.GetEvents(context.Background(), "AddCollateral", queryOptions)
	if err != nil {
		t.Fatal(err)
	}

	if len(collateralEvents) <= 0 {
		t.Logf("no events")
	}

	for i, e := range collateralEvents {
		t.Logf("%v) Event: %v. tx: %v", i, e.EventName, e.Transaction.TxHash)
	}
}
