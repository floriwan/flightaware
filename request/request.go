package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type flight struct {
	Ident       string
	FaFlightId  string
	ActualOff   string
	ActualOn    string
	Origin      airport
	Destination airport
}

type flights struct {
	Flights []flight
}

type airport struct {
	Code           string
	AirportInfoUrl string
}

const apiUrl = "https://aeroapi.flightaware.com/aeroapi"

func FlightInfo(reg string, apiKey string) flights {
	flights, err := search(apiUrl, reg, apiKey)
	if err != nil {
		log.Printf("%v", err)
	}
	return flights
}

// flights/search?query=-idents+AFL2381+-aboveAltitude+2

func search(url string, reg string, apiKey string) (flights, error) {
	url = url + "/flights/search?query=-idents+" + reg + "+-aboveAltitude+2"
	return sendHttpRequest(url, apiKey)
}

func sendHttpRequest(url string, apiKey string) (respF flights, err error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return flights{}, err
	}

	r.Header.Set("Accept", "application/json; charset=UTF-8")
	r.Header.Set("x-apikey", apiKey)

	log.Printf("> %v %v", r.Method, r.URL.Host+r.URL.RequestURI())

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return flights{}, err
	}
	defer resp.Body.Close()

	log.Printf("< status %v (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		return flights{}, fmt.Errorf("response status code was %v (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// read all data
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return flights{}, err
	}

	if len(b) == 0 {
		log.Printf("no data received")
		return flights{}, nil
	}

	// decode json response
	dec := json.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(&respF)
	if err != nil {
		return flights{}, fmt.Errorf("response '%v' %v", string(b), err)
	}

	return respF, nil
}
