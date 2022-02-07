package main

import (
	"flag"
	"fmt"

	"github.com/floriwan/flightaware/request"
)

const apiKey = ""

func main() {

	var dummy bool
	var reg string

	flag.BoolVar(&dummy, "d", false, "run in dummy mode, do not send any srequest")
	flag.StringVar(&reg, "r", "", "request information for aircraft registration")
	flag.Parse()

	flights := request.FlightInfo(reg, apiKey, dummy)

	fmt.Printf("response flights %+v", flights)
}
