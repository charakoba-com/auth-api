package authapi_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	authapi "github.com/charakoba-com/auth-api"
)

func TestHealthCheckHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(authapi.HealthCheckHandler))
	defer ts.Close()

	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(map[string]string{
		"message": "hello, world",
	})

	r, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if string(data) != buf.String() {
		t.Errorf("%s != %s", data, buf.String())
		return
	}
}

func TestAlgorithmHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(authapi.GetAlgorithmHandler))
	defer ts.Close()

	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(map[string]string{
		"alg": "RS512",
	})

	r, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if string(data) != buf.String() {
		t.Errorf("%s != %s", data, buf.String())
		return
	}
}
