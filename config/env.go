package config

import (
	"os"
	"strconv"
)

type Environment struct {
	Debug     bool
	APISecret string
}

var ENV = new(Environment)

func Refresh() {
	ENV.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	ENV.APISecret = os.Getenv("API_SECRET")
}
