package clients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const AlphaClientContextKey = ClientsContextKey("alpha")

type HttpAlphaClient struct {
}

type IAlphaClient interface {
	GetCats(ctx context.Context) []string
}

type CatsResponseType struct {
	Cats []string `json:"cats"`
}

func (HttpAlphaClient) GetCats(ctx context.Context) []string {
	httpclient := &http.Client{}
	res, err := httpclient.Get("http://localhost:9001/cats")
	if err != nil {
		panic("AARGH")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("AARGH")
	}
	catsresponse := &CatsResponseType{}
	err = json.Unmarshal(body, catsresponse)
	if err != nil {
		panic("AARGH")
	}
	return catsresponse.Cats
}

func ExtractFromContext(ctx context.Context) IAlphaClient {
	possibleAlphaClient := ctx.Value(AlphaClientContextKey)
	if possibleAlphaClient != nil {
		dependency, ok := possibleAlphaClient.(IAlphaClient)
		if ok {
			return dependency
		}
	}
	panic("myPrerogative could not be found in Request Context")
}
