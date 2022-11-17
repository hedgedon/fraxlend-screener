package main

import (
	"context"
	"encoding/json"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"log"
	"testing"
)

func GenerateReusableContracts(sdk *thirdweb.ThirdwebSDK) map[string]*thirdweb.SmartContract {
	fraxlendPairContracts := make(map[string]*thirdweb.SmartContract)
	for _, pair := range FraxlendPairList {
		contract, err := sdk.GetContractFromAbi(pair, FRAXLEND_PAIR_ABI)
		if err != nil {
			log.Fatal(err)
		}
		fraxlendPairContracts[pair] = contract
	}
	return fraxlendPairContracts
}

func TestCreatePairContracts(t *testing.T) {
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Fatal(err)
	}
	contracts := GenerateReusableContracts(sdk)
	resp, err := json.MarshalIndent(contracts, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Fraxlend Pair Contracts: %v", string(resp))
}

func TestPairDataContract(t *testing.T) {
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Fatal(err)
	}
	fraxlendPairContracts := GenerateReusableContracts(sdk)
	for pair, _ := range fraxlendPairContracts {
		name, err := fraxlendPairContracts[pair].Call(context.Background(), "name")
		if err != nil {
			t.Logf("Name err: %v", err)
		}
		symbol, err := fraxlendPairContracts[pair].Call(context.Background(), "symbol")
		if err != nil {
			t.Logf("Symbol err: %v", err)
		}
		pairAccounting, err := fraxlendPairContracts[pair].Call(context.Background(), "getPairAccounting")
		if err != nil {
			t.Logf("PairAccounting err: %v", err)
		}
		totalBorrow, err := fraxlendPairContracts[pair].Call(context.Background(), "totalBorrow")
		if err != nil {
			t.Logf("TotalBorrow err: %v", err)
		}
		totalAsset, err := fraxlendPairContracts[pair].Call(context.Background(), "totalAsset")
		if err != nil {
			t.Logf("TotalAsset err: %v", err)
		}
		totalSupply, err := fraxlendPairContracts[pair].Call(context.Background(), "totalSupply")
		if err != nil {
			t.Logf("TotalSupply err: %v", err)
		}
		maxLtv, err := fraxlendPairContracts[pair].Call(context.Background(), "maxLTV")
		if err != nil {
			t.Logf("MaxLTV err: %v", err)
		}
		t.Logf("Name: %v", name)
		t.Logf("PairAccounting: %v", pairAccounting)
		t.Logf("Symbol: %v", symbol)
		t.Logf("TotalBorrow: %v \n", totalBorrow)
		t.Logf("TotalAsset: %v \n", totalAsset)
		t.Logf("TotalSupply: %v \n", totalSupply)
		t.Logf("MaxLTV: %v \n", maxLtv)
		t.Logf(" ")
	}
}

func TestUserDataContract(t *testing.T) {
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Fatal(err)
	}
	randomUser := "0x3689c216f8f6ce7e2CE2a27c81a23096A787F532"
	contracts := GenerateReusableContracts(sdk)
	snapshot, err := contracts[FXS_FRAX_POOL].Call(context.Background(), "getUserSnapshot", randomUser)
	if err != nil {
		t.Logf("Snapshot err: %v", err)
	}
	userBorrowShares, err := contracts[FXS_FRAX_POOL].Call(context.Background(), "userBorrowShares", randomUser)
	if err != nil {
		t.Logf("UserBorrowShares err: %v", err)
	}
	userCollateralBalance, err := contracts[FXS_FRAX_POOL].Call(context.Background(), "userCollateralBalance", randomUser)
	if err != nil {
		t.Logf("UserBorrowShares err: %v", err)
	}
	t.Logf("Snapshot: %v", snapshot)
	t.Logf("UserBorrowShares: %v", userBorrowShares)
	t.Logf("UserCollateralBalance: %v", userCollateralBalance)
}
