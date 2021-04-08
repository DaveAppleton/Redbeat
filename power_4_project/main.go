package main

import (
	"html/template"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"af.sorry.gasprice/gasnow"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

// start OMIT
var (
	globalGasNowData struct {
		Rapid    uint64
		Fast     uint64
		Slow     uint64
		Standard uint64
	}
)

func gasLoop() {
	for {
		gnData, err := gasnow.GetLatest()
		if err != nil {
			log.Println("gasLoop", err)
			time.Sleep(1 * time.Minute)
			continue
		}
		atomic.StoreUint64(&globalGasNowData.Rapid, gnData.Rapid)
		atomic.StoreUint64(&globalGasNowData.Fast, gnData.Fast)
		atomic.StoreUint64(&globalGasNowData.Standard, gnData.Standard)
		atomic.StoreUint64(&globalGasNowData.Slow, gnData.Slow)
		time.Sleep(1 * time.Minute)
	}
}

// p1 OMIT
func gas(w http.ResponseWriter, r *http.Request) {
	var gnData struct {
		Rapid    uint64
		Fast     uint64
		Slow     uint64
		Standard uint64
	}
	gnData.Rapid = atomic.LoadUint64(&globalGasNowData.Rapid) / 1000000000
	gnData.Fast = atomic.LoadUint64(&globalGasNowData.Fast) / 1000000000
	gnData.Standard = atomic.LoadUint64(&globalGasNowData.Standard) / 1000000000
	gnData.Slow = atomic.LoadUint64(&globalGasNowData.Slow) / 1000000000
	// template_start OMIT
	//index, err := template.ParseFiles("files/main.html")
	index, err := template.Parse("files/main.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, gnData); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// template_end OMIT
}

// p2 OMIT
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

// p3 OMIT
func main() {
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	go gasLoop()
	http.HandleFunc("/", gas)
	http.ListenAndServe(":8080", nil)
}

// end OMIT
