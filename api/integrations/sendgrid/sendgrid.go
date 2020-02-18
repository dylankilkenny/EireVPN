package sendgrid

import (
	cfg "eirevpn/api/config"
	"eirevpn/api/models"
	"fmt"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGrid type
type SendGrid struct {
	Request rest.Request
}

// Send builds the request
func Send() *SendGrid {
	sg := SendGrid{}
	request := sendgrid.GetRequest(cfg.Load().SendGrid.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	sg.Request = request
	return &sg
}

func (sg *SendGrid) makeRequest() error {
	_, err := sendgrid.API(sg.Request)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

// RegistrationMail builds the body for sending a registration email
func (sg *SendGrid) RegistrationMail(user models.User, token string) error {
	m := mail.NewV3Mail()
	address := "info@eirevpn.ie"
	name := "ÉireVPN"
	e := mail.NewEmail(name, address)
	m.SetFrom(e)
	m.SetTemplateID(cfg.Load().SendGrid.Templates.Registration)
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(user.FirstName+" "+user.LastName, user.Email))
	p.SetDynamicTemplateData("confirm_email_url", cfg.Load().App.Domain+"/confirm_email?token="+token)
	m.AddPersonalizations(p)
	sg.Request.Body = mail.GetRequestBody(m)
	return sg.makeRequest()
}

// RegistrationMail builds the body for sending a registration email
func (sg *SendGrid) SupportRequest(email, subject, message string) error {
	m := mail.NewV3Mail()
	name := "Mail Service"
	e := mail.NewEmail(name, "mailservice@eirevpn.ie")
	m.SetFrom(e)
	m.SetTemplateID(cfg.Load().SendGrid.Templates.SupportRequest)
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail("Support", "support@eirevpn.ie"))
	p.AddCCs(mail.NewEmail("", email))
	p.Subject = subject
	p.SetDynamicTemplateData("subject", subject)
	p.SetDynamicTemplateData("message", message)
	m.AddPersonalizations(p)
	sg.Request.Body = mail.GetRequestBody(m)
	return sg.makeRequest()
}
