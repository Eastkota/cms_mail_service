package schema

import (
    "cms_mail_service/helpers"
    "cms_mail_service/model"

    "fmt"
    "context"
    "net/http"

    "github.com/graphql-go/graphql"
)

func AuthMiddleware(next func(p graphql.ResolveParams) *model.GenericMailResponse) func(p graphql.ResolveParams) *model.GenericMailResponse {
    return func(p graphql.ResolveParams) *model.GenericMailResponse {
        ctx := p.Context
        userInterface := ctx.Value("user")

        var user *model.User
        if userInterface == nil {
            if req, ok := ctx.Value("http_request").(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := helpers.ValidateToken(authHeader)
                if err != nil {
                    return helpers.FormatError(err)
                }
                if u == nil {
                    return helpers.FormatError(fmt.Errorf("invalid_token"))
                }

                ctx = context.WithValue(ctx, "user", u)
                p.Context = ctx
                user = u
            } else {
                return helpers.FormatError(fmt.Errorf("invalid_token"))
            }
        } else {
            user, _ = userInterface.(*model.User)
        }

        if user == nil {
            return helpers.FormatError(fmt.Errorf("invalid_token"))
        }
        return next(p)
    }
}