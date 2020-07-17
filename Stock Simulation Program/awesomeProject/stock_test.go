package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var High int
var Low int
var allStocks []StockStruct

func TestMain(t *testing.M) {
	High = 35
	Low = 5
	fmt.Println("High: ", High)
	fmt.Println("Low: ", Low)
	// initialize Stocks
	allStocks = make([]StockStruct, 0)
	allStocks = append(allStocks, StockStruct{Name: "aap", Value: 15})
	allStocks = append(allStocks, StockStruct{Name: "mcc", Value: 12})
	allStocks = append(allStocks, StockStruct{Name: "orc", Value: 17})
	allStocks = append(allStocks, StockStruct{Name: "thr", Value: 11})
	allStocks = append(allStocks, StockStruct{Name: "amz", Value: 16})
	allStocks = append(allStocks, StockStruct{Name: "eby", Value: 14})
	allStocks = append(allStocks, StockStruct{Name: "onz", Value: 18})
	allStocks = append(allStocks, StockStruct{Name: "ira", Value: 19})
	allStocks = append(allStocks, StockStruct{Name: "ipo", Value: 17})
	allStocks = append(allStocks, StockStruct{Name: "etf", Value: 10})
	allStocks = append(allStocks, StockStruct{Name: "apr", Value: 13})
	allStocks = append(allStocks, StockStruct{Name: "tsa", Value: 15})
	allStocks = append(allStocks, StockStruct{Name: "psp", Value: 18})
	allStocks = append(allStocks, StockStruct{Name: "nav", Value: 10})
	allStocks = append(allStocks, StockStruct{Name: "loi", Value: 11})
	allStocks = append(allStocks, StockStruct{Name: "gnp", Value: 17})
	allStocks = append(allStocks, StockStruct{Name: "frb", Value: 13})
	allStocks = append(allStocks, StockStruct{Name: "eft", Value: 13})
	allStocks = append(allStocks, StockStruct{Name: "arm", Value: 17})
	allStocks = append(allStocks, StockStruct{Name: "fdi", Value: 15})
	os.Exit(t.Run())
}

func TestStockSimulation(t *testing.T) {
	stopChan:=make(chan bool)
	go simulateStockPrices(allStocks, High, Low, stopChan)
	// Wait for 1-5 Minute then stop the simulation
	select {
	case <-time.After(time.Second*180):
		t.Log("Stopping channel")
		stopChan<-true
	}
}

func BenchmarkStockSimulation(b *testing.B) {
	stopChan:=make(chan bool)
	for n := 0; n < b.N; n++ {
		go simulateStockPrices(allStocks, High, Low, stopChan)
		// Wait for 1-5 Minute then stop the simulation
		select {
		case <-time.After(time.Second * 180):
			b.Log("Stopping channel")
			stopChan <- true
		}
	}
}
