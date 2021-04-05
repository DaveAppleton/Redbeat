// origin OMIT
package main

import (
	"fmt"
	"log"

	"af.sorry.power_1_test/grams2ounces"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
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

// mid OMIT
func main() {
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	weightInGrammes := viper.GetInt("WEIGHT_IN_GRAMMES")
	fmt.Println("weight in grammes", weightInGrammes)
	log.Println("weight in ounces", grams2ounces.GramsToOunces(weightInGrammes))
}

// end OMIT
