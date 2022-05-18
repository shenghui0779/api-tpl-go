package lib

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func Nonce() string {
	nonce := make([]byte, 8)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
}

func QueryPage(r *http.Request) (offset, limit int) {
	limit = 20

	if v, ok := URLQueryInt(r, "size"); ok && v > 0 {
		limit = int(v)
	}

	if limit > 100 {
		limit = 100
	}

	if v, ok := URLQueryInt(r, "page"); ok && v > 0 {
		offset = (int(v) - 1) * limit
	}

	return
}

func ExcelColumnIndex(name string) int {
	name = strings.ToUpper(name)

	if ok, err := regexp.MatchString(`^[A-Z]{1,2}$`, name); err != nil || !ok {
		return -1
	}

	index := 0

	for i, v := range name {
		if i != 0 {
			index = (index + 1) * 26
		}

		index += int(v - 'A')
	}

	return index
}
