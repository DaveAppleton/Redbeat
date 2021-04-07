package exdb

import (
	"context"
	"database/sql"
	"log"

	"time"

	"github.com/fsnotify/fsnotify"
	_ "github.com/lib/pq" // DB selection
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
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

func init() {
	initViper()
}

// start OMIT
type exchange struct {
	From  string
	To    string
	Price decimal.Decimal
}

func NewExchange(From, To string) (ex *exchange) {
	ex = &exchange{From: From, To: To}
	return
}

// mid OMIT
func (ex *exchange) UpdateExchange(price decimal.Decimal) (int, error) {
	if ex.Price == price {
		return 0, nil // nothing to do here
	}
	ex.Price = price
	var ID int
	dbConfig := viper.GetString("DB_CONFIG")
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Println("updateExchange 0: ", err)
		return 0, err
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sqlString := "insert into exchange (source,dest,price,date_sent) values($1,$2,$3,now()) returning id"
	err = db.QueryRowContext(ctx, sqlString, ex.From, ex.To, ex.Price).Scan(&ID)
	if err != nil {
		log.Println("updateExchange 1: ", err)
	}
	return ID, err
}

// end OMIT
