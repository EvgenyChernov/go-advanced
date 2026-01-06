package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinkCreateResponse struct {
	Hash string `json:"hash"`
}

type LinkAllLinksResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
