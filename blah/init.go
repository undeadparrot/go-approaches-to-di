package blah

import (
	"context"

	"github.com/undeadparrot/demoserver/wah"
)

func AddSmellToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, wah.MyKey("smell"), "blah")
}

func GetSmell(ctx context.Context) string {
	value := ctx.Value(wah.MyKey("smell"))
	if value != nil {
		return value.(string)
	}
	return "?"
}
