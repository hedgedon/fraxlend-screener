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

func TestFraxlendContractEvents_AddCollateral(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Errorf("error: %v \n", err)
	}
	contract, err := sdk.GetContractFromAbi(FXS_FRAX_POOL, FRAXLEND_PAIR_ABI)
	filters := map[string]interface{}{
		"_borrower": common.HexToAddress("0xd1F2739ad714045BE6146915275D0a2B822Ec1CC"),
	}
	queryOptions := thirdweb.EventQueryOptions{
		Filters: filters,
	}
	events, _ := contract.Events.GetEvents(context.Background(), "AddCollateral", queryOptions)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) <= 0 {
		t.Logf("no events")
	}
	for i, e := range events {
		t.Logf("%v) Event: %v. tx: %v", i, e.EventName, e.Transaction.TxHash)
	}
}

func TestFraxlendContractEvents_BorrowAsset(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Errorf("error: %v \n", err)
	}
	contract, err := sdk.GetContractFromAbi(FXS_FRAX_POOL, FRAXLEND_PAIR_ABI)
	filters := map[string]interface{}{
		"_borrower": common.HexToAddress("0xd1F2739ad714045BE6146915275D0a2B822Ec1CC"),
	}
	queryOptions := thirdweb.EventQueryOptions{
		Filters: filters,
	}
	events, _ := contract.Events.GetEvents(context.Background(), "BorrowAsset", queryOptions)
	if err != nil {
		panic(err)
	}
	for _, e := range events {
		t.Logf("Event: %v. tx: %v", e.EventName, e.Transaction.TxHash)
	}
}

// TestPairDataEvent: Filter a Fraxlend Pairs Event Occurrences
func TestPairDataEvent(t *testing.T) {
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Errorf("error: %v \n", err)
	}
	// If no filter needed:
	queryOptions := thirdweb.EventQueryOptions{}
	// Or Filter by specific User
	//randomUser := "0x3689c216f8f6ce7e2CE2a27c81a23096A787F532"
	//filters := map[string]interface{}{
	//	"_borrower": common.HexToAddress(randomUser),
	//}
	//queryOptions := thirdweb.EventQueryOptions{
	//	Filters: filters,
	//}
	contracts := GenerateReusableContracts(sdk)
	for pair, _ := range contracts {
		contract, err := sdk.GetContractFromAbi(pair, FRAXLEND_PAIR_ABI)
		if err != nil {
			t.Errorf("contract err: %v", err)
		}
		eventsAddCollateral, _ := contract.Events.GetEvents(context.Background(), "AddCollateral", queryOptions)
		if err != nil {
			t.Errorf("AddCollateral err: %v", err)
		}
		eventsBorrowAsset, _ := contract.Events.GetEvents(context.Background(), "BorrowAsset", queryOptions)
		if err != nil {
			t.Errorf("BorrowAsset err: %v", err)
		}
		t.Logf("Pair: %v", pair)
		t.Logf("# of AddCollateral events: %v", len(eventsAddCollateral))
		t.Logf("# of BorrowAsset events: %v", len(eventsBorrowAsset))
	}
}
