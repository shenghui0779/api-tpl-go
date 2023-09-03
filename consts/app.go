package consts

const MaxFormMemory = 32 << 20

type ContentType string

const (
	URLEncodedForm ContentType = "application/x-www-form-urlencoded"
	MultipartForm  ContentType = "multipart/form-data"
	ContentJSON    ContentType = "application/json"
)
