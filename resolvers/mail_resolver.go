package resolvers

import (
	"cms_mail_service/helpers"
	"cms_mail_service/model"
	"cms_mail_service/services"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

type MailResolver struct {
	Services services.Services
}

func NewEmailLogResolver(service services.Services) *MailResolver {
	return &MailResolver{Services: service}
}

func (r *MailResolver) FetchEmailLog(p graphql.ResolveParams) *model.GenericMailResponse {
    emails, err := r.Services.FetchEmailLog()
    if err != nil {
        return helpers.FormatError(err)
    }
    return &model.GenericMailResponse{
        Data: &model.MultipleMailData{
            EmailLog: emails,
        },
        Error: nil,
    }
}

func (r *MailResolver) SendMail(p graphql.ResolveParams) *model.GenericMailResponse {
	var mailInput model.MailInput
	inputData := p.Args["input"].(map[string]interface{})
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = json.Unmarshal(jsonData, &mailInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	err = r.Services.SendMail(mailInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericMailResponse{
		Data: &model.GenericMailSuccessData{
			Message: "Email sent succesfully",
		},
		Error: nil,
	}
}
