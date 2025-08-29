package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url string `json:"url"`
}

type LinksGetResponse struct {
	Links      []Link `json:"links"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int64  `json:"total"`
	TotalPages int64  `json:"total_pages"`
}

// type LinkGetResponse struct {
// 	Id        string `json:"id"`
// 	Url       string `json:"url"`
// 	Hash      string `json:"hash"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// }
