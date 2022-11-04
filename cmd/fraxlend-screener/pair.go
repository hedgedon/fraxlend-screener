package main

import (
	"fmt"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

func (s *Scope) createPairContracts() map[string]*thirdweb.SmartContract {
	fraxlendPairContracts := make(map[string]*thirdweb.SmartContract)
	for _, pair := range FraxlendPairList {
		contract, err := s.sdk.GetContractFromAbi(pair, FRAXLEND_PAIR_ABI)
		if err != nil {
			panic(err)
		}
		fraxlendPairContracts[pair] = contract
	}
	fmt.Println("created fraxlendPairContracts:", fraxlendPairContracts)
	return fraxlendPairContracts
}

func (s *Scope) getPairData() {
	name, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "name")
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", name)

	symbol, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "symbol")
	if err != nil {
		panic(err)
	}
	fmt.Println("symbol:", symbol)

	pairAccounting, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "getPairAccounting")
	if err != nil {
		panic(err)
	}
	for _, v := range pairAccounting.([]interface{}) { // use type assertion to loop over []interface{}
		fmt.Println(v)
	}
	fetchValue(pairAccounting)
	/*
		_totalAssetAmount
		Total assets deposited and interest accrued, total claims

		_totalAssetShares
		Total fTokens

		_totalBorrowAmount
		Total borrows

		_totalBorrowShares
		Total borrow shares

		_totalCollateral
		Total collateral
	*/

	totalBorrow, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "totalBorrow")
	if err != nil {
		panic(err)
	}
	fmt.Println("total borrow: ", totalBorrow)
	// total borrow:  [3521540165111577988854598 3508827027810177699518635]
	//TOTAL BORROW VALUE $3.52m

	totalAsset, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "totalAsset")
	if err != nil {
		panic(err)
	}
	fmt.Println("total Asset: ", totalAsset)

	totalSupply, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "totalSupply")
	if err != nil {
		panic(err)
	}
	fmt.Println("total Supply: ", totalSupply)

	maxLtv, err := s.contracts[FXS_FRAX_POOL].Call(s.ctx, "maxLTV")
	if err != nil {
		panic(err)
	}
	fmt.Println("max LTV:", maxLtv)
}
