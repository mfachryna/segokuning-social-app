package entity

type Post struct {
	ID         string   `json:"-"`
	UserId     string   `json:"-"`
	PostInHtml string   `json:"postInHtml"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"createdAt"`
}
