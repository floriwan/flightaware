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
const dummyResponse = `{"flights":[{"ident":"AFL2381","fa_flight_id":"AFL2381-1643694337-airline-0026","actual_off":"2022-02-03T12:00:53Z","actual_on":null,"predicted_out":null,"predicted_off":null,"predicted_on":null,"predicted_in":null,"predicted_out_source":null,"predicted_off_source":null,"predicted_on_source":null,"predicted_in_source":null,"origin":{"code":"LSGG","airport_info_url":"/airports/LSGG"},"destination":{"code":"UUEE","airport_info_url":"/airports/UUEE"},"waypoints":[],"first_position_time":"2022-02-03T11:45:03Z","last_position":{"fa_flight_id":"AFL2381-1643694337-airline-0026","altitude":11,"altitude_change":"D","groundspeed":141,"heading":75,"latitude":55.97455,"longitude":37.28622,"timestamp":"2022-02-03T15:16:49Z","update_type":"X"},"bounding_box":[56.10626,5.95486,46.08224,37.28622],"ident_prefix":null,"aircraft_type":"B738"}],"links":null,"num_pages":1}`

var dummyMode bool

func FlightInfo(reg string, apiKey string, dummy bool) flights {
	dummyMode = dummy
	flights, err := search(apiUrl, reg, apiKey)
	if err != nil {
		log.Printf("%v", err)
	}
	return flights
}

func search(url string, reg string, apiKey string) (respF flights, err error) {
	url = url + "/flights/search?query=-idents+" + reg + "+-aboveAltitude+2"
	if !dummyMode {
		return sendHttpRequest(url, apiKey)
	}

	dec := json.NewDecoder(bytes.NewReader([]byte(dummyResponse)))
	err = dec.Decode(&respF)
	if err != nil {
		return flights{}, fmt.Errorf("response '%v' %v", dummyResponse, err)
	}
	return respF, nil
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
