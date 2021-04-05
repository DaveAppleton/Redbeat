package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// start OMIT
var delayPeriod time.Duration

const dataForAPI = `{
	"result": "ok",
	"data": {
		"buy": 206.92,
		"sell": 201.18,
		"spot_price": 205.29542397091146,
		"timestamp": "2019-09-24T15:22:22.798+00:00"
	}
}`

func goldPrice(w http.ResponseWriter, r *http.Request) {
	time.Sleep(delayPeriod)
	w.Header().Set("Content_type", "application/json")
	fmt.Fprint(w, dataForAPI)
}

// end OMIT

func setDelay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_type", "text/plain")
	delayStr := r.FormValue("delay")
	if delayStr == "" {
		fmt.Fprintln(w, "Please set delay parameter")
		return
	}
	delayP, err := strconv.ParseInt(delayStr, 10, 64)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	delayPeriod = time.Second * time.Duration(delayP)
	fmt.Fprintln(w, "duration set to ", delayPeriod)
}

func main() {
	http.HandleFunc("/gold_price", goldPrice)
	http.HandleFunc("/delay", setDelay)
	http.ListenAndServe(":8080", nil)
}

// fin OMIT
