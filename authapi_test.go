package authapi_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	authapi "github.com/charakoba-com/auth-api"
	"github.com/charakoba-com/auth-api/db"
)

const testPassword = "testpasswd"

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

func TestPostUserHandler(t *testing.T) {
	db.Init(nil)
	testUsername := "createuser"
	ts := httptest.NewServer(http.HandlerFunc(authapi.PostUserHandler))
	defer ts.Close()

	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(map[string]interface{}{
		"message": "SUCCESS",
	})

	payload := bytes.Buffer{}
	payload.WriteString(`{"username":"`)
	payload.WriteString(testUsername)
	payload.WriteString(`", "password":"testpasswd"}`)

	r, err := http.Post(ts.URL, "application/json", &payload)
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
	t.Logf(string(data))
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	user := db.User{}
	if err = user.Lookup(tx, testUsername); err != nil {
		t.Errorf("%s", err)
		return
	}

	if user.Name != testUsername {
		t.Errorf("%s != %s", user.Name, testUsername)
		return
	}
	if err = user.Delete(tx); err != nil {
		t.Errorf("%s", err)
		return
	}
}
