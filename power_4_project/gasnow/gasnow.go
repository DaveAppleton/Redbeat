package gasnow

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	//"golang.org/x/net/websocket"
)

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

// gasNowData OMIT
type gasNowData struct {
	Rapid     uint64
	Fast      uint64
	Slow      uint64
	Standard  uint64
	Timestamp int64
}

// gl_start OMIT
// GetLatest gas prices from gasnow.org
func GetLatest() (gnData *gasNowData, err error) {
	var gnResp struct {
		Code int
		Data gasNowData
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	gasURL := "https://www.gasnow.org/api/v3/gas/price?utm_source=RedBeat"

	req, err := http.NewRequest("GET", gasURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// gl_mid OMIT

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("bad status code :" + resp.Status)
	}
	//gnData = new(gasNowData)
	if err = json.NewDecoder(resp.Body).Decode(&gnResp); err != nil {
		return nil, err
	}
	return &(gnResp.Data), nil
}

// gl_end OMIT
