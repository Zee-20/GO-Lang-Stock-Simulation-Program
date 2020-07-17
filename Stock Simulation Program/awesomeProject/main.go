package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

type StockStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	High := 35
	Low := 5
	fmt.Println("High: ", High)
	fmt.Println("Low: ", Low)
	// initialize Stocks
	all_stocks := []StockStruct{}
	all_stocks = append(all_stocks, StockStruct{Name: "aap", Value: 15})
	all_stocks = append(all_stocks, StockStruct{Name: "mcc", Value: 12})
	all_stocks = append(all_stocks, StockStruct{Name: "orc", Value: 17})
	all_stocks = append(all_stocks, StockStruct{Name: "thr", Value: 11})
	all_stocks = append(all_stocks, StockStruct{Name: "amz", Value: 16})
	all_stocks = append(all_stocks, StockStruct{Name: "eby", Value: 14})
	all_stocks = append(all_stocks, StockStruct{Name: "onz", Value: 18})
	all_stocks = append(all_stocks, StockStruct{Name: "ira", Value: 19})
	all_stocks = append(all_stocks, StockStruct{Name: "ipo", Value: 17})
	all_stocks = append(all_stocks, StockStruct{Name: "etf", Value: 10})
	all_stocks = append(all_stocks, StockStruct{Name: "apr", Value: 13})
	all_stocks = append(all_stocks, StockStruct{Name: "tsa", Value: 15})
	all_stocks = append(all_stocks, StockStruct{Name: "psp", Value: 18})
	all_stocks = append(all_stocks, StockStruct{Name: "nav", Value: 10})
	all_stocks = append(all_stocks, StockStruct{Name: "loi", Value: 11})
	all_stocks = append(all_stocks, StockStruct{Name: "gnp", Value: 17})
	all_stocks = append(all_stocks, StockStruct{Name: "frb", Value: 13})
	all_stocks = append(all_stocks, StockStruct{Name: "eft", Value: 13})
	all_stocks = append(all_stocks, StockStruct{Name: "arm", Value: 17})
	all_stocks = append(all_stocks, StockStruct{Name: "fdi", Value: 15})
	fmt.Println("Initials Stocks Values")
	fmt.Println(all_stocks)
	stopChan := make(chan bool)
	go simulateStockPrices(all_stocks, High, Low, stopChan)
	http.ListenAndServe(":8080",nil)
}

func simulateStockPrices(all_stocks []StockStruct, High int, Low int, controller chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-controller:
			fmt.Printf("Stopping stock simulator")
			return
		case t, _ := <-ticker.C:
			fmt.Println("\nTime: ", t)

			for i := 0; i <= 19; i++ {
				randon_value := rand.Intn(26)
				inc_or_dec_random := rand.Intn(2)
				if inc_or_dec_random == 0 {
					fmt.Println("Increment  by value: ", randon_value, " ", all_stocks[i].Name)
					updateStocks(all_stocks, i, randon_value, "inc")

				} else {
					fmt.Println("Decrement whole stock by value: ", randon_value, " ", all_stocks[i].Name)
					updateStocks(all_stocks, i, randon_value, "dec")
				}

				if all_stocks[i].Value >= High {
					text := fmt.Sprintf("Bought! Stock Name: %s Stock Value: %d", all_stocks[i].Name, all_stocks[i].Value)
					fmt.Println(text)
					writeTofile(text)
				}

				if all_stocks[i].Value <= Low {
					text := fmt.Sprintf("Sold! Stock Name: %s Stock Value: %d", all_stocks[i].Name, all_stocks[i].Value)
					fmt.Println(text)
					writeTofile(text)
				}
			}
			fmt.Println("\nUpadated Stocks: ")
			fmt.Println(all_stocks)
		}
	}
}

func updateStocks(arr []StockStruct, index int, value int, op string) {
	if len(arr) > 0 {
		if op == "inc" {
			arr[index].Value = arr[index].Value + value
		}
		if op == "dec" {
			arr[index].Value = arr[index].Value - value
		}
	}
}

func writeTofile(text string) {
	f, err := os.OpenFile("stocks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text + "\n"); err != nil {
		fmt.Println(err)
	}
}

// func RandStringBytes(n int) string {
// 	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = letterBytes[rand.Intn(len(letterBytes))]
// 	}
// 	return string(b)
// }
