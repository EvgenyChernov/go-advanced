package link

import (
	"app/adv-http/pkg/request"
	"app/adv-http/pkg/response"
	"fmt"
	"net/http"
)

type LinkHendlerDeps struct {
	LinkRepository *LinkRepository
}
type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHendlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /link/{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, req)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		Createdlink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.JsonResponse(w, Createdlink, 200)
	}
}

func (handler *LinkHandler) GoTo() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, req, link.URL, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Println(id)

	}
}

func (handler *LinkHandler) Update() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
