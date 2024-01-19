package internal

import (
	"encoding/json"
	"io"
	"net/http"

	yiigo_http "github.com/shenghui0779/yiigo/http"
	yiigo_util "github.com/shenghui0779/yiigo/util"
	"github.com/shenghui0779/yiigo/validator"
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
	switch yiigo_util.ContentType(r) {
	case yiigo_http.ContentForm:
		if err := r.ParseForm(); err != nil {
			return err
		}
	case yiigo_http.ContentFormMultipart:
		if err := r.ParseMultipartForm(yiigo_http.MaxFormMemory); err != nil {
			if err != http.ErrNotMultipart {
				return err
			}
		}
	}

	if err := yiigo_util.MapForm(obj, r.Form); err != nil {
		return err
	}

	return validator.ValidateStruct(obj)
}
