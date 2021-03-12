package response

type BookInfo struct {
	Title       string  `json:"title"`
	SubTitle    string  `json:"subtitle"`
	Author      string  `json:"author"`
	Version     string  `json:"version"`
	Price       float64 `json:"price"`
	Publisher   string  `json:"publisher"`
	PublishDate string  `json:"publish_date"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}
