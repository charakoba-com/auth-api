package authapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/charakoba-com/auth-api/db"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	HealthCheckMessage string
	Algorithm          string
}

var config Config

func init() {
	log.Printf("Initialize API Config...")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.authapi")
	viper.AddConfigPath(".")
	viper.SetDefault("healthCheckMessage", "hello, world")
	viper.SetDefault("algorithm", "RS512")
	viper.ReadInConfig()
	config = Config{
		HealthCheckMessage: viper.GetString("healthCheckMessage"),
		Algorithm:          viper.GetString("algorithm"),
	}
}

// ANY /
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": config.HealthCheckMessage,
	}
	httpJSON(w, response)
}

// POST /user
// create user handler
func postUserHandler(w http.ResponseWriter, r *http.Request) {
	u := db.User{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpError(w, 400, `reading request`, err)
		return
	}

	var createUserRequest db.CreateUserRequest
	if err = json.Unmarshal(data, &createUserRequest); err != nil {
		httpError(w, 500, `invalid json`, err)
		return
	}

	u.Name = createUserRequest.Name
	u.Password = createUserRequest.Password
	u.CreatedOn = time.Now()
	u.ModifiedOn = time.Now()
	log.Printf(`CREATE USER :: {"username": "%s", "password": "%s"}`, u.Name, u.Password)
	tx, err := db.BeginTx()
	if err != nil {
		httpError(w, 500, `database transaction`, err)
		return
	}

	if err = u.Create(tx); err != nil {
		httpError(w, 500, `creating user`, err)
		return
	}
	response := map[string]string{
		"message": "SUCCESS",
	}
	httpJSON(w, response)
}

// DELETE /user
// delete user handler
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
}

// GET /user/list
// get user list handler
func getUserListHandler(w http.ResponseWriter, r *http.Request) {
}

// POST /auth
// authentication handler
func postAuthHandler(w http.ResponseWriter, r *http.Request) {
}

// GET /algorithm
// get using algorithm handler
func getAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"alg": config.Algorithm,
	}
	httpJSON(w, response)
}

// POST /verify
// verification token handler
func postVerifyHandler(w http.ResponseWriter, r *http.Request) {
}

// GET /key
// get public key handler
func getKeyHandler(w http.ResponseWriter, r *http.Request) {
}
