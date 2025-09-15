package schema

import "github.com/graphql-go/graphql"

var EmailLogType = graphql.NewObject(graphql.ObjectConfig{
    Name: "EmailLog",
    Fields: graphql.Fields{
        "id":            &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
        "receiver":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "subject":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "title":         &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "content":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "status":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "error_message": &graphql.Field{Type: graphql.String},
        "created_at":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
        "updated_at":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
    },
})

var MultipleMailData = graphql.NewObject(graphql.ObjectConfig{
	Name: "MultipleMailData",
	Fields: graphql.Fields{
		"email_log": &graphql.Field{Type: graphql.NewList(EmailLogType)},
	},
})
