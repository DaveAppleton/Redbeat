package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shopspring/decimal"
)

// start OMIT
const goldPriceServer = "http://127.0.0.1:8080/gold_price"

type spotData struct {
	Buy       decimal.Decimal
	Sell      decimal.Decimal
	SpotPrice decimal.Decimal `json:"spot_price"`
	Timestamp time.Time
}

type spotReply struct {
	Result string
	Data   spotData
}

var spot spotReply

// end OMIT

// start code OMIT

func main() {
	req, err := http.NewRequest("GET", goldPriceServer, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&spot)
	if err != err {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(spot.Data.Buy, spot.Data.Sell, spot.Data.SpotPrice, spot.Data.Timestamp)
}

// end code OMIT
