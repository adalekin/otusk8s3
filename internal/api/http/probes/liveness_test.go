package probes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiveness(t *testing.T) {
	w := httptest.NewRecorder()
	LivenessHandler(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	status, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if have, want := string(status), "{\"status\":\"OK\"}"; have != want {
		t.Errorf("The status is wrong. Have: %s, want: %s.", have, want)
	}
}
