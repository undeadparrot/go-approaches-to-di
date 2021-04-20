package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type myContextKey string

const nameKey = myContextKey("name")
const envKey = myContextKey("env")

func get_name(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, nameKey, r.URL.Path[1:])
	x0 := ctx.Value(nameKey).(string)
	x1 := ctx.Value(envKey).(string)

	fmt.Fprintf(w, "Hi there, I love %s! (%s)", x0, x1)
}

func _main() {
	env, hasEnv := os.LookupEnv("ENV")
	if !hasEnv {
		env = "LOCAL"
	}
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), envKey, env)
			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}

	main_with_gorilla_mux(middleware)
}

func main_with_gorilla_mux(middleware mux.MiddlewareFunc) {
	r := mux.NewRouter()
	r.Use(middleware)
	r.HandleFunc("/{name}", get_name)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func _main_with_http(middleware mux.MiddlewareFunc) {

	http.Handle("/", middleware(http.HandlerFunc(get_name)))
	fmt.Println("Hello World! :)")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func main() {

	ctx := context.WithValue(context.Background(), envKey, "TEST")

	req, err := http.NewRequestWithContext(ctx, "GET", "/Scott", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(get_name)
	handler.ServeHTTP(rr, req)

	if !strings.Contains(rr.Body.String(), "Scott") {
		log.Fatal("wrong!")
	}
	log.Print("Success!")
}
