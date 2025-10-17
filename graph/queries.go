package schema

import (
	"cms_mail_service/helpers"
	"cms_mail_service/model"
	"cms_mail_service/resolvers"

	"github.com/graphql-go/graphql"
)


func NewQueryType(resolver *resolvers.MailResolver) *graphql.Object {
    return graphql.NewObject(graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "service": &graphql.Field{
                Type: graphql.NewNonNull(Service),
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    schema, err := GetSchema()
                    if err != nil {
                        return nil, err
                    }
                    serviceInfo := model.Service{
                        Name:    "mailservice",
                        Version: "1.0.0",
                        Schema:  helpers.ConvertSchemaToString(schema),
                    }
                    return serviceInfo, nil
                },
            },
            "fetchEmailLog": &graphql.Field{
                Type: EmailLogResponse,
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    return AuthMiddleware(PermissionMiddleware("list", resolver.FetchEmailLog))(p), nil
                },
            },
        },
    })
}