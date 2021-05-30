package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shenghui0779/yiigo"
)

var validator = yiigo.NewValidator()

func BindJSON(r *http.Request, obj interface{}) error {
	defer io.Copy(io.Discard, r.Body)
	
	if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
		return err
	}

	if err := validator.ValidateStruct(obj); err != nil {
		return err
	}

	return nil
}
