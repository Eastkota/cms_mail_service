package schema

import "github.com/graphql-go/graphql"

var EmailLogResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmailLogResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: MultipleMailData},
		"error": &graphql.Field{Type: MailError},
	},
})
var GenericMailResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericMailResponse",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: GenericMailSuccessData},
		"error": &graphql.Field{Type: MailError},
	},
})
