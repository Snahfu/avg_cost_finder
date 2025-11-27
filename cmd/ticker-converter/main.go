package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticker struct {
	FundTickerDateId int     `json:"FundTickerDateId"`
	TickerId         int     `json:"TickerId"`
	TickerCode       string  `json:"TickerCode"`
	Quantity         int     `json:"Quantity"`
	QtyAvailable     int     `json:"QtyAvailable"`
	UnitPrice        float64 `json:"UnitPrice"`
	CostPrice        float64 `json:"CostPrice"`
}

type TickerData struct {
	FundDateId  int      `json:"FundDateId"`
	FundId      int      `json:"FundId"`
	FundCode    string   `json:"FundCode"`
	StrategyID  int      `json:"StrategyID"`
	FundDate    string   `json:"FundDate"`
	FundStock   int      `json:"FundStock"`
	FundDeposit int      `json:"FundDeposit"`
	FundBond    int      `json:"FundBond"`
	FundCash    int      `json:"FundCash"`
	FundAUM     int      `json:"FundAUM"`
	FundAR      int      `json:"FundAR"`
	FundAP      int      `json:"FundAP"`
	Ticker      []Ticker `json:"Ticker"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/ticker-converter/main.go <input_file>")
		fmt.Println("Example: go run cmd/ticker-converter/main.go data.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var tickers []Ticker
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) != 7 {
			fmt.Printf("Skipping invalid line: %s\n", line)
			continue
		}

		fundTickerDateId, _ := strconv.Atoi(parts[0])
		tickerCode := parts[1]
		tickerId, _ := strconv.Atoi(parts[2])
		quantity, _ := strconv.Atoi(parts[3])
		qtyAvailable, _ := strconv.Atoi(parts[4])
		unitPrice, _ := strconv.ParseFloat(parts[5], 64)
		costPrice, _ := strconv.ParseFloat(parts[6], 64)

		ticker := Ticker{
			FundTickerDateId: fundTickerDateId,
			TickerId:         tickerId,
			TickerCode:       tickerCode,
			Quantity:         quantity,
			QtyAvailable:     qtyAvailable,
			UnitPrice:        unitPrice,
			CostPrice:        costPrice,
		}

		tickers = append(tickers, ticker)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create output structure
	currentDate := time.Now().Format("2006-01-02")
	output := TickerData{
		FundDateId:  0,
		FundId:      12,
		FundCode:    "XXX",
		StrategyID:  1,
		FundDate:    currentDate,
		FundStock:   0,
		FundDeposit: 0,
		FundBond:    0,
		FundCash:    0,
		FundAUM:     0,
		FundAR:      0,
		FundAP:      0,
		Ticker:      tickers,
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		os.Exit(1)
	}

	// Generate output filename with current date
	outputFile := fmt.Sprintf("ticker_%s.json", currentDate)

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %d tickers to %s\n", len(tickers), outputFile)
}
