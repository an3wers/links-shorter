package auth

import (
	"go/links-shorter/configs"
	"go/links-shorter/pkg/jwt"
	"go/links-shorter/pkg/req"
	"go/links-shorter/pkg/resp"
	"net/http"
	// "regexp"
)

type AuthHandlerDeps struct {
	DbConfig    *configs.DbConfig
	AuthConfig  *configs.AuthConfig
	AuthService *AuthService
}

type AuthHandler struct {
	DbConfig    *configs.DbConfig
	AuthConfig  *configs.AuthConfig
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		DbConfig:    deps.DbConfig,
		AuthConfig:  deps.AuthConfig,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginRequest](w, r)
		_ = body

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			resp.Json(w, http.StatusUnauthorized, err.Error())
			return
		}

		token, err := jwt.NewJWT(handler.AuthConfig.SecretKey).Create(jwt.JWTData{
			Email: user.Email,
		})

		if err != nil {
			resp.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := LoginResponse{
			Token: token,
		}

		resp.Json(w, http.StatusOK, data)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		_ = body

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		user, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		token, err := jwt.NewJWT(handler.AuthConfig.SecretKey).Create(jwt.JWTData{
			Email: user.Email,
		})

		if err != nil {
			resp.Json(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := RegisterResponse{
			Token: token,
		}

		resp.Json(w, http.StatusCreated, data)

	}
}
