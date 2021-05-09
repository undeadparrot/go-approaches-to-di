package tests

import (
	"context"
)

type DummyAlphaClient struct {
	GetCatsResponse []string
	PostCatResponse bool
}

func (d DummyAlphaClient) GetCats(ctx context.Context) []string {
	return d.GetCatsResponse
}
func (d DummyAlphaClient) PostCat(ctx context.Context, cat string) bool {
	return d.PostCatResponse
}

func defaultDummyAlphaClient() DummyAlphaClient {
	return DummyAlphaClient{
		GetCatsResponse: []string{"pussy meow"},
		PostCatResponse: true,
	}
}
