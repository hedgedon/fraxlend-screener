package main

import (
	"encoding/json"
	"fmt"
)

func fetchValue(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Printf("%v is an interface \n ", value)
	case bool:
		fmt.Printf("%v is bool \n ", value)
	case float64:
		fmt.Printf("%v is float64 \n ", value)
	case []interface{}:
		fmt.Printf("%v is a slice of interface \n ", value)
		for _, v := range value.([]interface{}) { // use type assertion to loop over []interface{}
			//fetchValue(v)
			fmt.Println(v)
			//return v
		}
	case map[string]interface{}:
		fmt.Printf("%v is a map \n ", value)
		for _, v := range value.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			fetchValue(v)
		}
	default:
		fmt.Printf("%v is unknown \n ", value)
	}
}

func pretty(o map[string]string) {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

//func pretty2(o map[string]*any) {
//	b, err := json.MarshalIndent(o, "", "  ")
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//	fmt.Print(string(b))
//}
