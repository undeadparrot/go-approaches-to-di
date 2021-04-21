package wah

import (
	"context"
)

type MyKey string

func GetSmell(ctx context.Context) string {
	value := ctx.Value(MyKey("smell"))
	if value != nil {
		return value.(string)
	}
	return "?"
}
