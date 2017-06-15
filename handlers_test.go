package authapi_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"

	authapi "github.com/charakoba-com/auth-api"
	"github.com/charakoba-com/auth-api/db"
	"github.com/charakoba-com/auth-api/model"
	"github.com/charakoba-com/auth-api/service"
)

var s *authapi.Server
var ts *httptest.Server

func TestMain(m *testing.M) {
	s = authapi.New()
	ts = httptest.NewServer(s)
	defer ts.Close()
	db.Init(nil)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestHealthCheckHandlerOK(t *testing.T) {
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer res.Body.Close()
	// I/O test
	var hcres model.HealthCheckResponse
	if err := json.NewDecoder(res.Body).Decode(&hcres); err != nil {
		t.Errorf("%s", err)
		return
	}
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}
	if hcres.Message != `hello, world` {
		t.Errorf(`"%s" != "hello, world"`, hcres.Message)
		return
	}
}

func TestHealthCheckHandlerMethodNotAllowed(t *testing.T) {
	requestBody := bytes.Buffer{}
	res, err := http.Post(ts.URL, "", &requestBody)
	// I/O test
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if res.StatusCode != 405 {
		t.Errorf("status 405 Method Not Allowed is expected, but %s", res.Status)
		return
	}
}

func TestCreateUserHandlerOK(t *testing.T) {
	path := "/user"
	t.Logf("POST %s", path)
	requestBody := bytes.Buffer{}
	requestBody.WriteString(`{"id": "createID", "username": "createdUser", "password": "testpasswd"}`)
	res, err := http.Post(ts.URL+path, "application/json", &requestBody)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}
	var createUserResponse model.CreateUserResponse
	if err := json.NewDecoder(res.Body).Decode(&createUserResponse); err != nil {
		t.Errorf("%s", err)
		return
	}
	if createUserResponse.Message != "success" {
		t.Errorf("response message is invalid")
		return
	}

	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer func() {
		var usrSvc service.UserService
		// reset
		if err := usrSvc.Delete(tx, "createID"); err != nil {
			t.Errorf("%s", err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Errorf("%s", err)
			return
		}
	}()

	// data test
	var usrSvc service.UserService
	user, err := usrSvc.Lookup(tx, `createID`)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	expectedUser := model.User{
		ID:       "createID",
		Name:     "createdUser",
		Password: "testpasswd",
	}
	if *user != expectedUser {
		t.Errorf("%s != %s", user, expectedUser)
		return
	}
}

func TestLookupUserHandlerOK(t *testing.T) {
	path := "/user/lookupID"
	t.Logf("GET %s", path)
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}

	var lures model.LookupUserResponse
	if err := json.NewDecoder(res.Body).Decode(&lures); err != nil {
		t.Errorf("%s", err)
		return
	}
	expectedUser := model.User{
		ID:   "lookupID",
		Name: "lookupuser",
	}
	if lures.User != expectedUser {
		t.Errorf("%s != %s", lures.User, expectedUser)
		return
	}
}

func TestLookupUserHandlerNotFound(t *testing.T) {
	path := "/user/hoge"
	t.Logf("GET %s", path)
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 404 {
		t.Errorf("status 404 Not Found is expected, but %s", res.Status)
		return
	}
	var errorResponse model.ErrorResponse
	if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
		t.Errorf("%s", err)
		return
	}
	if errorResponse.Message != "user not found" {
		t.Errorf(`"%s" != "user not found"`, errorResponse.Message)
		return
	}
}

func TestUpdateUserHandlerOK(t *testing.T) {
	path := "/user"
	t.Logf("PUT %s", path)
	updateUserRequest := model.UpdateUserRequest{
		ID:       "updateID",
		Username: "updateduser",
		Password: "testpasswd",
	}
	// I/O test
	requestBody := bytes.Buffer{}
	if err := json.NewEncoder(&requestBody).Encode(updateUserRequest); err != nil {
		t.Errorf("%s", err)
		return
	}
	req, err := http.NewRequest("PUT", ts.URL+path, &requestBody)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}
	var updateUserResponse model.UpdateUserResponse
	if err := json.NewDecoder(res.Body).Decode(&updateUserResponse); err != nil {
		t.Errorf("response is invalid json")
		return
	}
	if updateUserResponse.Message != "success" {
		t.Errorf("response message is invalid")
		return
	}
	// data test
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	var usrSvc service.UserService
	user, err := usrSvc.Lookup(tx, `updateID`)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	expectedUser := model.User{
		ID:       "updateID",
		Name:     "updateduser",
		Password: "testpasswd",
	}
	if *user != expectedUser {
		t.Errorf("%s != %s", user, expectedUser)
		return
	}
	// reset
	if err := usrSvc.Update(tx, &db.User{ID: "updateID", Name: "updateuser", Password: "testpasswd"}); err != nil {
		t.Errorf("%s", err)
		return
	}
	if err := tx.Commit(); err != nil {
		t.Errorf("%s", err)
		return
	}
}

func TestDeleteUserHandlerOK(t *testing.T) {
	path := "/user/deleteID"
	t.Logf("DELETE %s", path)
	req, err := http.NewRequest("DELETE", ts.URL+path, nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}
	var deleteUserResponse model.DeleteUserResponse
	if err := json.NewDecoder(res.Body).Decode(&deleteUserResponse); err != nil {
		t.Errorf("response is invalid json")
		return
	}
	if deleteUserResponse.Message != "success" {
		t.Errorf("response message is invalid")
		return
	}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// data test
	var usrSvc service.UserService
	_, err = usrSvc.Lookup(tx, `deleteID`)
	if err == nil {
		t.Errorf("sql.ErrNoRows should be occured, but there is no error")
		return
	}
	// reset
	if err := usrSvc.Create(tx, &db.User{ID: "deleteID", Name: "deleteuser", Password: "testpasswd"}); err != nil {
		t.Errorf("%s", err)
		return
	}
	if err := tx.Commit(); err != nil {
		t.Errorf("%s", err)
		return
	}
}

func TestListupUserHandlerOK(t *testing.T) {
	path := "/user/list"
	t.Logf("GET %s", path)
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 200 {
		t.Errorf("status 200 OK is expected, but %s", res.Status)
		return
	}
	expected := model.ListupUserResponse{
		Users: model.UserList{
			model.User{
				ID:   "lookupID",
				Name: "lookupuser",
			},
			model.User{
				ID:   "updateID",
				Name: "updateuser",
			},
			model.User{
				ID:   "deleteID",
				Name: "deleteuser",
			},
		},
	}
	var listupUserResponse model.ListupUserResponse
	if err := json.NewDecoder(res.Body).Decode(&listupUserResponse); err != nil {
		t.Errorf("%s", err)
		return
	}
	sort.Sort(expected.Users)
	sort.Sort(listupUserResponse.Users)
	for i, user := range listupUserResponse.Users {
		if user != expected.Users[i] {
			t.Errorf("%s != %s", user, expected.Users[i])
			return
		}
	}
}

func TestAuthHandlerOK(t *testing.T) {
}

func TestGetAlgorithmHandlerOK(t *testing.T) {
	urls := []string{
		ts.URL + "/algorithm",
		ts.URL + "/alg",
	}
	for _, url := range urls {
		t.Logf("%s", url)
		res, err := http.Get(url)
		if err != nil {
			t.Errorf("%s", err)
			return
		}
		defer res.Body.Close()
		// I/O test
		var gares model.GetAlgorithmResponse
		if err := json.NewDecoder(res.Body).Decode(&gares); err != nil {
			t.Errorf("%s", err)
			return
		}
		if gares.Algorithm != `RS256` {
			t.Errorf(`"%s" != "RS256"`, gares.Algorithm)
			return
		}
	}
}

func TestGetAlgorithmHandlerMethodNotAllowed(t *testing.T) {
	requestBody := bytes.Buffer{}
	res, err := http.Post(ts.URL, "", &requestBody)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	// I/O test
	if res.StatusCode != 405 {
		t.Errorf("status 405 Method Not Allowed is expected, but %s", res.Status)
		return
	}
}

func VerifyHandlerOK(t *testing.T) {
}

func GetKeyHandlerOK(t *testing.T) {
}
