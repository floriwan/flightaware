package main

import (
	"fmt"

	"github.com/floriwan/flightaware/request"
)

const apiKey = ""

func main() {

	flights := request.FlightInfo("AFL2381", apiKey)

	fmt.Printf("response flights %+v", flights)
}
