package consts

const MaxFormMemory = 32 << 20

type ContentType string

const (
	MIMEForm          ContentType = "application/x-www-form-urlencoded"
	MIMEMultipartForm ContentType = "multipart/form-data"
	ContentJSON       ContentType = "application/json"
)
