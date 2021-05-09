package handlers

import (
	"fmt"
	"net/http"

	"github.com/undeadparrot/demoserver/clients"
)

type myContextKey string

const NameKey = myContextKey("name")
const EnvKey = myContextKey("env")

type NameHandler struct {
	SecretWord    string
	MyAlphaClient clients.IAlphaClient
}

func (h *NameHandler) GetNameHandler(w http.ResponseWriter, r *http.Request) {
	x0 := r.URL.Path[1:]

	cats := h.MyAlphaClient.GetCats(r.Context())

	fmt.Fprintf(w, "Hi there, I love %s! (%s) (%s)", x0, cats, h.SecretWord)
}

func (h *NameHandler) TradeCatsHandler(w http.ResponseWriter, r *http.Request) {
	x0 := r.URL.Path[1:]

	cats := h.MyAlphaClient.GetCats(r.Context())
	ok := h.MyAlphaClient.PostCat(r.Context(), "Kot")

	if ok {
		fmt.Fprintf(w, "Hi there, I love %s! (%s) (%s)", x0, cats, h.SecretWord)
	}
}
