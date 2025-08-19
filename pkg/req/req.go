package req

import (
	"go/links-shorter/pkg/resp"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := Decode[T](r.Body)

	if err != nil {
		resp.Json(w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	err = IsValid(body)

	if err != nil {
		resp.Json(w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	return &body, nil

}
