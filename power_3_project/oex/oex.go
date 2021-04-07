package oex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

// start OMIT
func initViper() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("viper config changed", e.Name)
	})
}

func init() {
	initViper()
}

// p1 OMIT
type RateResponse struct {
	Error     ResponseError
	Base      string
	Timestamp int
	Rates     map[string]decimal.Decimal
}

type ResponseError struct {
	Error       bool
	Status      int
	Message     string
	Description string
}

// p2 OMIT

func GetLatest() (response *RateResponse, err error) {
	apiKey := viper.GetString("RATE_ID")
	if len(apiKey) == 0 {
		return nil, errors.New("oex:GetLatest : RATE_ID not set")
	}

	response = new(RateResponse)
	response.Rates = make(map[string]decimal.Decimal)

	oex_latest := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", apiKey)
	req, err := http.NewRequest("GET", oex_latest, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//p3 OMIT
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var RE ResponseError
		if err = json.NewDecoder(resp.Body).Decode(&RE); err != nil {
			return nil, err
		}
		err = errors.New(RE.Message)
		response.Error = RE
		return

	}
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	return
}

// end OMIT
