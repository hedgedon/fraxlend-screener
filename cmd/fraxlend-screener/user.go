package main

import (
	"fmt"
)

/*
calculating LTV: https://docs.frax.finance/fraxlend/advanced-concepts/position-health-and-liquidations


borrow share price = totalBorrow.amount / totalBorrow.shares
LTV = (Borrow Shares x Share Price) / ( Collateral Balance / Exchange Rate )
(10,015.9144349971 * 1) / (3311.78 * 6.25) = 0.48
*/

// IF LTV is below X value, send me an alert on Discord
// monitor LTV position
// should be able to manage and monitor my Fraxlend position

func (s *Scope) getUserData(fraxlendPair, address string) {
	fmt.Println("")
	fmt.Println(">>>>> USER SNAPSHOT <<<<<<")
	// GET USER SNAPSHOT
	snapshot, err := s.contracts[fraxlendPair].Call(s.ctx, "getUserSnapshot", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("snapshot:", snapshot)

	userBorrowShares, err := s.contracts[fraxlendPair].Call(s.ctx, "userBorrowShares", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("userBorrowShares:", userBorrowShares)

	userCollateralBalance, err := s.contracts[fraxlendPair].Call(s.ctx, "userCollateralBalance", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("userCollateralBalance:", userCollateralBalance) // FXS

	//fxs := s.tokenList["FXS"]
	//tokenContract, err := s.sdk.GetToken(fxs)
	//if err != nil {
	//	panic(err)
	//}
	//token, err := tokenContract.BalanceOf(os.Getenv("MY_ADDRESS"))
	//if err != nil {
	//	panic(err)
	//}
	//balance := token.DisplayValue
	//fmt.Println("balance:", balance)
}
