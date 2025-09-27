package schema

import "github.com/graphql-go/graphql"


var MailInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MailInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"to": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"subject": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)