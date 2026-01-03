package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkCreateResponse struct {
	Hash string `json:"hash"`
}
