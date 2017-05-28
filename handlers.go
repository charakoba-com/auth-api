package authapi

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	HealthCheckMessage string
}

var config Config

func init() {
	log.Printf("Initialize API Config...")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.authapi")
	viper.AddConfigPath(".")
	viper.SetDefault("healthCheckMessage", "hello, world")
	viper.ReadInConfig()
	config = Config{
		HealthCheckMessage: viper.GetString("healthCheckMessage"),
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
}

// POST /verify
// verification token handler
func postVerifyHandler(w http.ResponseWriter, r *http.Request) {
}

// GET /key
// get public key handler
func getKeyHandler(w http.ResponseWriter, r *http.Request) {
}
