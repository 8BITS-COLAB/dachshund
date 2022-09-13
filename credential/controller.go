package credential

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CredentialController struct {
	cr *CredentialRepo
	d  *EmailService
}

func NewController(cr *CredentialRepo, d *EmailService) *CredentialController {
	return &CredentialController{
		cr: cr,
		d:  d,
	}
}

func (cc *CredentialController) Register(ctx *fiber.Ctx) error {
	var d CreateCredentialDTO

	if err := ctx.BodyParser(&d); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := d.Validate(); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	c := NewCredential(d)

	go cc.d.SendWelcome(c)

	if err := cc.cr.Create(*c); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(http.StatusCreated)
}

func (cc *CredentialController) SendEmail(ctx *fiber.Ctx) error {
	var d SendEmailDTO

	if err := ctx.BodyParser(&d); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := d.Validate(); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	c := ctx.Locals("credential").(*Credential)

	go cc.d.SendDefault(c, &d)

	return ctx.SendStatus(http.StatusOK)
}
