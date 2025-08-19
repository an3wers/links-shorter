package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type LinkDeleteRequest struct {
	Id string `json:"id"`
}

type LinkGetRequest struct {
	Id string `json:"id"`
}

type LinkGetResponse struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	Hash      string `json:"hash"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
