package internal

import (
	"encoding/json"
	"io"
	"net/http"

	libhttp "api/lib/http"
	"api/lib/util"
	"api/lib/validator"
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
	switch util.ContentType(r) {
	case libhttp.ContentForm:
		if err := r.ParseForm(); err != nil {
			return err
		}
	case libhttp.MultipartForm:
		if err := r.ParseMultipartForm(libhttp.MaxFormMemory); err != nil {
			if err != http.ErrNotMultipart {
				return err
			}
		}
	}

	if err := util.MapForm(obj, r.Form); err != nil {
		return err
	}

	return validator.ValidateStruct(obj)
}
