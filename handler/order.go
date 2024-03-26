package handler

import (
	"encoding/json"
	"fmt"
	"github.com/pindarEng/stockSimulator/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("created an account")
}
func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updated an account")
}
func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleted an account")
}
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read an account")
}
func (o *Order) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("read an account")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing any necessary data (nil in this case)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type TimeSeries struct {
	MetaData   map[string]string            `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series (5min)"`
}

func StockHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement api
	apiKey := "4P88T5KPRHW21GAI"
	symbol := "IBM"
	res, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + symbol + "&interval=5min&apikey=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stocks models.Stock
	err = json.Unmarshal(body, &stocks)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("stocks: %+v", stocks)

	latestTimeStamp := stocks.MetaData.LastRefreshed
	latestData, ok := stocks.TimeSeries[latestTimeStamp]
	if !ok {
		log.Fatalf("no data available:%s", latestTimeStamp)
	}

	fmt.Println("Symbol:", stocks.MetaData.Symbol)
	fmt.Println("Timestamp:", latestTimeStamp)
	fmt.Println("Open:", latestData.Open)
	fmt.Println("High:", latestData.High)
	fmt.Println("Low:", latestData.Low)
	fmt.Println("Close:", latestData.Close)
	fmt.Println("Volume:", latestData.Volume)

}
