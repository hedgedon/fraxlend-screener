package main

import (
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"testing"
)

// TestGetBalance: Displays the balance of token
func TestGetBalance(t *testing.T) {
	sdk, err := thirdweb.NewThirdwebSDK("mainnet", nil)
	if err != nil {
		t.Fatal(err)
	}
	tokenContract, err := sdk.GetToken(FXS)
	token, err := tokenContract.BalanceOf(FXS_FRAX_POOL)
	if err != nil {
		t.Fatal(err)
	}
	balance := token.DisplayValue
	token2Contract, err := sdk.GetToken(FRAX_TOKEN)
	if err != nil {
		t.Fatal(err)
	}
	token2, err := token2Contract.BalanceOf(FXS_FRAX_POOL)
	if err != nil {
		t.Fatal(err)
	}
	balance2 := token2.DisplayValue
	t.Logf("%v & %v", balance, balance2)
}
