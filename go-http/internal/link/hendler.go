package link

import (
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
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// handler.LinkRepository.Create(&Link{
		// 	URL:  req.URL.String(),
		// 	Hash: req.URL.Path,
		// })
	}
}

func (handler *LinkHandler) GoTo() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

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
