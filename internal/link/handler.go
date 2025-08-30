package link

import (
	"go/links-shorter/configs"
	"go/links-shorter/internal/stat"
	"go/links-shorter/pkg/middleware"
	"go/links-shorter/pkg/req"
	"go/links-shorter/pkg/resp"
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	Repo     *LinkRepository
	StatRepo *stat.StatRepository
	Config   *configs.Config
}

type LinkHandler struct {
	Repo     *LinkRepository
	StatRepo *stat.StatRepository
	Config   *configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {

	handler := &LinkHandler{
		Repo:     deps.Repo,
		StatRepo: deps.StatRepo,
	}

	// public
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("GET /{hash}", handler.GoTo())

	// auth
	router.Handle("GET /link/{id}", middleware.Auth(handler.GetLink(), deps.Config))
	router.Handle("GET /link", middleware.Auth(handler.GetLinks(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.Auth(handler.UpdateLink(), deps.Config))
	router.Handle("DELETE /link/{id}", middleware.Auth(handler.DeleteLink(), deps.Config))
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

		// fmt.Println("Context", r.Context().Value(middleware.ContextAuthKey).(string))

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
		page, err := strconv.Atoi(r.URL.Query().Get("page"))

		if err != nil {
			page = 1
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

		if err != nil {
			limit = 10
		}

		offset := (page - 1) * limit

		links, err := handler.Repo.GetLinks(limit, offset)
		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		if len(links) == 0 {
			links = []Link{}
		}

		total, err := handler.Repo.Count()

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		totalPages := int64(math.Ceil(float64(total) / float64(limit)))

		response := LinksGetResponse{
			Links:      links,
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		}

		resp.Json(w, http.StatusOK, response)
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

		handler.StatRepo.AddClick(link.ID)

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
