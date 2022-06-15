package config

import (
	"os"
	"strconv"
)

type Environment struct {
	Debug     bool
	ApiSecret string
}

var ENV = new(Environment)

func RefreshENV() {
	ENV.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	ENV.ApiSecret = os.Getenv("API_SECRET")
}
