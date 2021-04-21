package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/undeadparrot/demoserver/clients"
	"github.com/undeadparrot/demoserver/handlers"
)

func main() {
	ctx1 := context.WithValue(context.Background(), handlers.NameKey, "blah")
	val1 := ctx1.Value(handlers.NameKey)
	val2 := ctx1.Value("name")
	fmt.Printf("Context values %s %s", val1, val2)
	env, hasEnv := os.LookupEnv("ENV")
	if !hasEnv {
		env = "LOCAL"
	}
	injectEnvMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), handlers.EnvKey, env)
			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}

	alphaClient := &clients.HttpAlphaClient{}

	injectClientsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx = context.WithValue(ctx, clients.AlphaClientContextKey, alphaClient)

			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}

	r := mux.NewRouter()
	r.Use(injectEnvMiddleware, injectClientsMiddleware)
	r.HandleFunc("/{name}", handlers.GetNameHandler)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
