package config

import (
	"os"
	"strconv"
)

type Environment struct {
	Debug bool
}

var ENV = new(Environment)

func RefreshENV() {
	ENV.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
}
