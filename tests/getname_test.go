package tests

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/undeadparrot/demoserver/clients"
	"github.com/undeadparrot/demoserver/handlers"
)

type DummyClient struct{}

func (DummyClient) GetCats(ctx context.Context) []string {
	cats := []string{"Pussy meow"}
	return cats
}

func TestGetName(t *testing.T) {

	dummyClient := &DummyClient{}

	ctx := context.Background()
	// ctx = context.WithValue(ctx, handlers.EnvKey, "TEST")
	ctx = context.WithValue(ctx, clients.AlphaClientContextKey, dummyClient)

	req, err := http.NewRequestWithContext(ctx, "GET", "/Scott", nil)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetNameHandler)
	handler.ServeHTTP(rr, req)

	if !strings.Contains(rr.Body.String(), "Scott") {
		log.Fatal("wrong!")
	}
	log.Print("Success!")
}
