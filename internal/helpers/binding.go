package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/shenghui0779/yiigo"
)

var validator = yiigo.NewValidator()

func BindJSON(r *http.Request, obj interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return err
	}

	if err := validator.ValidateStruct(obj); err != nil {
		return err
	}

	return nil
}
