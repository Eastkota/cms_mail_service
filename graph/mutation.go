package schema

import (
	"cms_mail_service/resolvers"
	
	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.MailResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"sendMail": &graphql.Field{
				Type: GenericMailResponse,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: MailInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolver.SendMail(p), nil
				},
			},
		},
	})
}
