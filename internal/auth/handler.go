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

		// var payload LoginRequest
		// err := json.NewDecoder(r.Body).Decode(&payload)

		// if err != nil {
		// 	resp.Json(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// validate := validator.New()
		// err = validate.Struct(payload)

		if err != nil {
			resp.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		// // check empty email
		// // check empty password
		// if payload.Email == "" || payload.Password == "" {
		// 	resp.Json(w, http.StatusBadRequest, "email or password is empty")
		// 	return
		// }

		// // check is valid email
		// reg, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		// if !reg.MatchString(payload.Email) {
		// 	resp.Json(w, http.StatusBadRequest, "email is not valid")
		// 	return
		// }

		// // check is min length password
		// if len(payload.Password) < 8 {
		// 	resp.Json(w, http.StatusBadRequest, "password is not valid")
		// 	return
		// }

		user, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			resp.Json(w, http.StatusUnauthorized, err.Error())
			return
		}

		token, err := jwt.NewJWT(handler.AuthConfig.SecretKey).Create(user.Email)

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

		token, err := jwt.NewJWT(handler.AuthConfig.SecretKey).Create(user.Email)

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
