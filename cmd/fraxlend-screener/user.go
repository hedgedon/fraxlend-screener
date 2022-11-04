package main

import (
	"fmt"
)

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
	fmt.Println("userCollateralBalance:", userCollateralBalance)
}
