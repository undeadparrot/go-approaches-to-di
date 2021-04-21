package handlers

import (
	"fmt"
	"net/http"

	"github.com/undeadparrot/demoserver/clients"
)

type myContextKey string

const NameKey = myContextKey("name")
const EnvKey = myContextKey("env")

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	x0 := r.URL.Path[1:]

	alphaclient := clients.ExtractFromContext(r.Context())
	cats := alphaclient.GetCats(r.Context())

	fmt.Fprintf(w, "Hi there, I love %s! (%s)", x0, cats)
}
