package link

import (
	"go/links-shorter/pkg/req"
	"go/links-shorter/pkg/resp"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	Repo *LinkRepository
}

type LinkHandler struct {
	Repo *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {

	handler := &LinkHandler{
		Repo: deps.Repo,
	}

	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("GET /link/{id}", handler.GetLink())
	router.HandleFunc("GET /link", handler.GetLinks())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("PATCH /link/{id}", handler.UpdateLink())
	router.HandleFunc("DELETE /link/{id}", handler.DeleteLink())
}

func (handler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](w, r)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}
		link := NewLink(body.Url)
		for {
			existed, _ := handler.Repo.GetByHash(link.Hash)
			if existed == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.Repo.CreateLink(link)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		resp.Json(w, http.StatusCreated, createdLink)

	}
}

func (handler *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](w, r)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		updatedLink, err := handler.Repo.UpdateLink(&Link{
			Model: gorm.Model{
				ID: uint(id),
			},
			Url: body.Url,
		})

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		resp.Json(w, http.StatusOK, updatedLink)
	}
}

func (handler *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		link, err := handler.Repo.GetLinkById(uint(id))

		if err != nil {
			resp.Json(w, http.StatusNotFound, err.Error())
			return
		}

		resp.Json(w, http.StatusOK, link)

	}
}

func (handler *LinkHandler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := handler.Repo.GetByHash(hash)

		if err != nil {
			resp.Json(w, http.StatusNotFound, err.Error())
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}

func (handler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = handler.Repo.GetLinkById(uint(id))

		if err != nil {
			resp.Json(w, http.StatusNotFound, err.Error())
			return
		}

		err = handler.Repo.DeleteById(uint(id))

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		resp.Json(w, http.StatusOK, struct{}{})

	}
}
