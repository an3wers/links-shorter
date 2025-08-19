package link

import (
	"go/links-shorter/pkg/req"
	"go/links-shorter/pkg/resp"
	"net/http"
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
	router.HandleFunc("PUT /link/{id}", handler.UpdateLink())
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

	}
}

func (handler *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// fmt.Println(id)

		resp.Json(w, http.StatusOK, struct {
			Id string `json:"id"`
		}{id})
	}
}
