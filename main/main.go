package main

import (
	"flag"
	"fmt"

	"github.com/floriwan/flightaware/request"
)

const apiKey = "xx"

func main() {

	var dummy bool

	flag.BoolVar(&dummy, "d", false, "run in dummy mode, do not send any srequest")

	flights := request.FlightInfo("AFL2381", apiKey, dummy)

	fmt.Printf("response flights %+v", flights)
}
