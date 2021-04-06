// origin OMIT
package main

import (
	"fmt"
	"log"
	"os"

	"af.sorry.power_1_test/grams2ounces"
	"github.com/fsnotify/fsnotify"
	"github.com/ogier/pflag"
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
	var weightInGrammes int
	pflag.IntVar(&weightInGrammes, "weight_in_grammes", -1, "weight to convert to ounces")
	pflag.Parse()
	if weightInGrammes == -1 {
		pflag.Usage()
		os.Exit(1)
	}
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})

	fmt.Println("weight in grammes", weightInGrammes)
	fmt.Println("weight in ounces", grams2ounces.GramsToOunces(weightInGrammes))
	log.Println("weight in grammes", weightInGrammes)
	log.Println("weight in ounces", grams2ounces.GramsToOunces(weightInGrammes))
}

// end OMIT
