package main

import (
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

// end OMIT

// start code OMIT
var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func main() {
	var spot spotReply
	resp, err := httpClient.Get(goldPriceServer)
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
