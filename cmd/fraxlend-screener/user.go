package main

import (
	"fmt"
	"os"
)

func (s *Scope) getUserData() {
	MyAddress := os.Getenv("MY_ADDRESS")

	// GET USER SNAPSHOT
	snapshot, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "getUserSnapshot", MyAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("snapshot:", snapshot)

	userBorrowShares, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "userBorrowShares", MyAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("userBorrowShares:", userBorrowShares)

	userCollateralBalance, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "userCollateralBalance", MyAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("userCollateralBalance:", userCollateralBalance)
}
