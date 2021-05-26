package base

import "encoding/json"

type PostBody interface {
	Bytes() ([]byte, error)
}

type jsonbody struct {
	data interface{}
}

func (j *jsonbody) Bytes() ([]byte, error) {
	return json.Marshal(j.data)
}

func NewJSONBody(params interface{}) PostBody {
	return &jsonbody{
		data: params,
	}
}
