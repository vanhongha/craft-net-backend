package directives

import (
	"context"
	"craftnet/internal/app/middleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	if lo.IsNil(tokenData) {
		return nil, &gqlerror.Error{
			Message:    "Access Denied",
			Extensions: map[string]interface{}{},
		}
	}

	return next(ctx)
}
