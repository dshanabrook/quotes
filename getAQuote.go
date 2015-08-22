package main

import (
	"fmt"
	"os"

	"github.com/stocks"
)

var ticker = "AAPL"

func main() {
	ticker := os.Args[1]
	stockFromYahoo, err := stocks.GetQuote(ticker)
	if err != nil {
		fmt.Println(err)
	}
	aQuote, err := stockFromYahoo.GetPrice()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("this a a quote %6.2f of a value\n", aQuote)
}
