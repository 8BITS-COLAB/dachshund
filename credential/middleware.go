package credential

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewMiddleware(cr *CredentialRepo) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		apiKey := ctx.Get("dachshund-api-key")

		if apiKey == "" {
			return fiber.NewError(http.StatusUnauthorized, "unauthorized")
		}

		c := cr.GetByApiKey(apiKey)

		if c == nil {
			return fiber.NewError(http.StatusUnauthorized, "unauthorized")
		}

		ctx.Locals("credential", c)

		return ctx.Next()
	}
}
