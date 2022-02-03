package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testResp = `{"flights":[{"ident":"AFL2381","fa_flight_id":"AFL2381-1643694337-airline-0026","actual_off":"2022-02-03T12:00:53Z","actual_on":null,"predicted_out":null,"predicted_off":null,"predicted_on":null,"predicted_in":null,"predicted_out_source":null,"predicted_off_source":null,"predicted_on_source":null,"predicted_in_source":null,"origin":{"code":"LSGG","airport_info_url":"/airports/LSGG"},"destination":{"code":"UUEE","airport_info_url":"/airports/UUEE"},"waypoints":[],"first_position_time":"2022-02-03T11:45:03Z","last_position":{"fa_flight_id":"AFL2381-1643694337-airline-0026","altitude":11,"altitude_change":"D","groundspeed":141,"heading":75,"latitude":55.97455,"longitude":37.28622,"timestamp":"2022-02-03T15:16:49Z","update_type":"X"},"bounding_box":[56.10626,5.95486,46.08224,37.28622],"ident_prefix":null,"aircraft_type":"B738"}],"links":null,"num_pages":1}`

func TestReqeust(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")
		w.Write([]byte(testResp))
	}))
	defer ts.Close()

	f, err := search(ts.URL, "", "")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", f)
	if len(f.Flights) < 1 {
		t.Errorf("expected search response")
	}
}
