package request

import (
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

type airport struct {
	Code           string
	AirportInfoUrl string
}

const apiUrl = "https://aeroapi.flightaware.com/aeroapi/"

func FlightInfo(reg string, apiKey string) {
	search(apiUrl, reg, apiKey)
}

// flights/search?query=-idents+AFL2381+-aboveAltitude+2

func search(url string, reg string, apiKey string) error {
	url = url + "flights/search?query=-idents+" + reg + "+-aboveAltitude+2"
	return sendHttpRequest(url, apiKey)
}

func sendHttpRequest(url string, apiKey string) error {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	r.Header.Set("Accept", "application/json; charset=UTF-8")
	r.Header.Set("x-apikey", apiKey)

	log.Printf("> %v %v", r.Method, r.URL.Host+r.URL.RequestURI())

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("< status %v (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))

	return nil
}
