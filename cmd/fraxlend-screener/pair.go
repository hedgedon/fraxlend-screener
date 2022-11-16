package main

import (
	"fmt"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func (s *Scope) createPairContracts(contractAddresses []string) map[string]*thirdweb.SmartContract {
	fraxlendPairContracts := make(map[string]*thirdweb.SmartContract)
	for _, pair := range contractAddresses {
		contract, err := s.sdk.GetContractFromAbi(pair, FRAXLEND_PAIR_ABI)
		if err != nil {
			panic(err)
		}
		fraxlendPairContracts[pair] = contract
	}
	fmt.Println("created mapping of Fraxlend Pair Contract:", fraxlendPairContracts)
	return fraxlendPairContracts
}

func (s *Scope) getPairData(fraxlendPair string) {
	fmt.Println("")
	name, err := s.contracts[fraxlendPair].Call(s.ctx, "name")
	if err != nil {
		panic(err)
	}

	symbol, err := s.contracts[fraxlendPair].Call(s.ctx, "symbol")
	if err != nil {
		panic(err)
	}

	pairAccounting, err := s.contracts[fraxlendPair].Call(s.ctx, "getPairAccounting")
	if err != nil {
		panic(err)
	}
	//for _, v := range pairAccounting.([]interface{}) { // use type assertion to loop over []interface{}
	//	fmt.Println(v)
	//}
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

	totalBorrow, err := s.contracts[fraxlendPair].Call(s.ctx, "totalBorrow")
	if err != nil {
		panic(err)
	}
	//fmt.Println("total borrow: ", totalBorrow)
	// total borrow:  [3521540165111577988854598 3508827027810177699518635]
	//TOTAL BORROW VALUE $3.52m

	totalAsset, err := s.contracts[fraxlendPair].Call(s.ctx, "totalAsset")
	if err != nil {
		panic(err)
	}

	totalSupply, err := s.contracts[fraxlendPair].Call(s.ctx, "totalSupply")
	if err != nil {
		panic(err)
	}

	maxLtv, err := s.contracts[fraxlendPair].Call(s.ctx, "maxLTV")
	if err != nil {
		panic(err)
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf(">> %v (%v)", name, symbol)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Printf("Max LTV: %v \n", maxLtv)
	fmt.Println("total borrow: ", totalBorrow)
	fmt.Println("total Asset: ", totalAsset)
	fmt.Println("total Supply: ", totalSupply)
	fmt.Println("pair accounting:")
	fetchValue(pairAccounting)
}
