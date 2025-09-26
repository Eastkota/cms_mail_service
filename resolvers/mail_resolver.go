package resolvers

import (
	"cms_mail_service/helpers"
	"cms_mail_service/model"
	"cms_mail_service/services"

	"github.com/graphql-go/graphql"
)

type MailResolver struct {
	Services services.Services
}

func NewEmailLogResolver(service services.Services) *MailResolver {
	return &MailResolver{Services: service}
}

func (r *MailResolver) FetchEmailLog(p graphql.ResolveParams) *model.EmailLogResponse {
    emails, err := r.Services.FetchEmailLog()
    if err != nil {
        return helpers.FormatError(err)
    }
    return &model.EmailLogResponse{
        Data: &model.MultipleMailData{
            EmailLog: emails,
        },
        Error: nil,
    }
}
