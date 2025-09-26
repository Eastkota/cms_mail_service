package resolvers

import (
	"cms_mail_service/model"

	"github.com/graphql-go/graphql"
)

type Resolver interface {
	FetchEmailLog(p graphql.ResolveParams) ([]model.GenericMailResponse, error)
}
