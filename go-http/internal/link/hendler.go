package link

import (
	"app/adv-http/configs"
	"app/adv-http/pkg/middleware"
	"app/adv-http/pkg/request"
	"app/adv-http/pkg/response"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHendlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}
type LinkHandler struct {
	LinkRepository *LinkRepository
}

func (handler *LinkHandler) GoAll() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		limitInt, err := strconv.Atoi(req.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		offsetInt, err := strconv.Atoi(req.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		links := handler.LinkRepository.GetAll(limitInt, offsetInt)
		count := handler.LinkRepository.Count()
		response.JsonResponse(w, LinkAllLinksResponse{
			Links: links,
			Count: count,
		}, 200)
	}
}

func NewLinkHandler(router *http.ServeMux, deps LinkHendlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.Handle("DELETE /link/{id}", middleware.IsAuthenticated(handler.Delete(), deps.Config))
	router.HandleFunc("GET /link/{hash}", handler.GoTo())
	router.HandleFunc("GET /link", handler.GoAll())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, req)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existingLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existingLink == nil {
				break
			}
			link.GenerateHash()
		}
		Createdlink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.JsonResponse(w, Createdlink, 200)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
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

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.JsonResponse(w, nil, 200)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		email, ok := req.Context().Value(middleware.ContextKeyEmail).(string)
		if ok {
			fmt.Println(email)
		}

		body, err := request.HandleBody[LinkUpdateRequest](w, req)
		if err != nil {
			return
		}
		idString := req.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(Link{
			Model: gorm.Model{ID: uint(id)},
			URL:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.JsonResponse(w, link, 200)
	}
}
