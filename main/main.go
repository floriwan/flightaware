package main

import (
	"fmt"

	"github.com/floriwan/flightaware/request"
)

const apiKey = "UcV4NKtlrBGAsqRmzWuQDMVfATriybCh"

func main() {

	flights := request.FlightInfo("AFL2381", apiKey)

	fmt.Printf("response flights %+v", flights)
}
