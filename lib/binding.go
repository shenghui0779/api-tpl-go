package lib

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shenghui0779/yiigo"
	"github.com/shenghui0779/yiigo/validator"
	"github.com/shenghui0779/yiigo/xhttp"
)

func BindJSON(r *http.Request, obj any) error {
	if r.Body != nil && r.Body != http.NoBody {
		defer io.Copy(io.Discard, r.Body)

		if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
			return err
		}
	}

	return validator.ValidateStruct(obj)
}

// BindForm 解析Form表单并校验
func BindForm(r *http.Request, obj any) error {
	switch yiigo.ContentType(r) {
	case xhttp.ContentForm:
		if err := r.ParseForm(); err != nil {
			return err
		}
	case xhttp.ContentFormMultipart:
		if err := r.ParseMultipartForm(xhttp.MaxFormMemory); err != nil {
			if err != http.ErrNotMultipart {
				return err
			}
		}
	}

	if err := yiigo.MapForm(obj, r.Form); err != nil {
		return err
	}

	return validator.ValidateStruct(obj)
}
