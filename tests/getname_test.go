package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/undeadparrot/demoserver/handlers"
)

func TestGetName(t *testing.T) {

	req, err := http.NewRequest("GET", "/Scott", nil)
	if err != nil {
		log.Fatal(err)
	}

	dummyClient := defaultDummyAlphaClient()
	dummyClient.GetCatsResponse = []string{"puss puss", "mew mew"}
	h := &handlers.NameHandler{SecretWord: "Blah", MyAlphaClient: dummyClient}

	handler := http.HandlerFunc(h.GetNameHandler)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if !strings.Contains(rr.Body.String(), "Scott") {
		log.Fatal("wrong!")
	}
	log.Print("Success!")
}
