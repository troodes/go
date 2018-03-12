package main

// Some good resources:
// https://github.com/GoClipse/goclipse/blob/latest/documentation/UserGuide.md#user-guide
// https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
// https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/

import (
	"fmt"
	"github.com/troodes/sandbox/utils"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Starting the application...")

	// Read config.json input file.
	var myConfig = utils.LoadConfiguration("./hello/config.json")
	var ticker string = "DIS"
	var data string

	// Example of hitting google datasource
	fmt.Println("================= Google Data =====================")
	var url string = "https://finance.google.com/finance/getprices?p=2d&f=c&q=" + ticker
	data = getDataFromUrl(url)
	fmt.Println(data)

	// Example of hitting AlphaVantage Data Source
	fmt.Println("================= AlphaVantage Data =====================")
	url = "https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=" + ticker + "&apikey=" + myConfig.AvKey
	fmt.Println(url)
	// Example of hitting WorldTradingData Data Source
	data = getDataFromUrl(url)
	fmt.Println(data)

	fmt.Println("================= WorldTradingData Data =====================")
	url = "https://www.worldtradingdata.com/api/v1/stock?api_token=" + myConfig.WtdKey + "&symbol=" + ticker
	fmt.Println(url)
	data = getDataFromUrl(url)
	fmt.Println(data)

	fmt.Println("Terminating the application...")
}

func getDataFromUrl(url string) string {

	var data string
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}

	return string(data)
}