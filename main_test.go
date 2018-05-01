package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	//URL
	BASE_URL := "https://www.cryptopia.co.nz/api/GetCurrencies"

	var market marketCurrency

	requestDataFromAPI(BASE_URL, &market)
	showCurrenciesList(&market)

	var btc Data

	btc = market.getCurrency("BTC")
	fmt.Print("Algoritmo -> ")
	fmt.Println(btc.Algorithm)
}

func TestMarket(t *testing.T) {
	//URL
	BASE_URL := "https://www.cryptopia.co.nz/api/GetMarket/DOT_BTC"
	var market MarketTrade

	//request data
	response, err := http.Get(BASE_URL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&market); err != nil {
		log.Println(err)
	}
	fmt.Printf("%f ", market.Data.AskPrice)

}

func requestDataFromAPI(url string, market *marketCurrency) {
	//request data
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&market); err != nil {
		log.Println(err)
	}
}

func showCurrenciesList(market *marketCurrency) {

	sortByID(market)

	for index := range market.Data {
		fmt.Printf("%d %s\n", market.Data[index].ID, market.Data[index].Name)

	}
}

func sortByID(market *marketCurrency) {
	for i := 0; i < len(market.Data); i++ {
		dataAux := market.Data[i]
		for j := i - 1; j >= 0 && dataAux.ID < market.Data[j].ID; j-- {
			market.Data[j+1] = market.Data[j]
			market.Data[j] = dataAux
		}
	}
}
