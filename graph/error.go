package schema

import "github.com/graphql-go/graphql"

var MailError = graphql.NewObject(graphql.ObjectConfig{
	Name: "MailError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
