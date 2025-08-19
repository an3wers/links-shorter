package halthcheck

import (
	"fmt"
	"net/http"
)

type HalthHandler struct{}

func NewHalthHandler(router *http.ServeMux) {
	handler := &HalthHandler{}
	router.HandleFunc("GET /halthcheck", handler.Halth())
}

func (handler *HalthHandler) Halth() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("halthcheck")
	}
}
