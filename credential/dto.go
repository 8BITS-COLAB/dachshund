package credential

import (
	"errors"
	"net/mail"
)

type CreateCredentialDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (d *CreateCredentialDTO) Validate() error {
	if len(d.Name) < 3 {
		return errors.New("name must be min 3 characters")
	}

	if _, err := mail.ParseAddress(d.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}

type ContextBody struct {
	Subject     string `json:"subject"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Context struct {
	To   *Credential `json:"to"`
	Body ContextBody `json:"body"`
}

type SendEmailDTO struct {
	Template string  `json:"template"`
	Context  Context `json:"context"`
}

func (d *SendEmailDTO) Validate() error {
	if d.Template == "" {
		return errors.New("template is required")
	}

	if len(d.Context.To.Name) < 3 {
		return errors.New("to name must be min 3 characters")
	}

	if _, err := mail.ParseAddress(d.Context.To.Email); err != nil {
		return errors.New("invalid to email")
	}

	if len(d.Context.Body.Subject) < 3 {
		return errors.New("subject must be min 3 characters")
	}

	if len(d.Context.Body.Title) < 3 {
		return errors.New("body title must be min 3 characters")
	}

	return nil
}
