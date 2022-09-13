package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ElioenaiFerrari/dachshund/credential"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), nil)

	if err != nil {
		log.Fatal(err)
	}

	// conn.Migrator().DropTable("credentials")
	cr := credential.NewRepo(conn)
	d := credential.NewEmailService()

	cc := credential.NewController(cr, d)

	conn.AutoMigrate(&credential.Credential{})

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := http.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			return ctx.Status(code).JSON(fiber.Map{"message": err.Error()})
		},
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowMethods: "POST,GET"}))
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(compress.New())

	app.Static("/img", "asset/img")
	app.Static("/css", "asset/css")

	v1 := app.Group("/api/v1")
	app.Get("/_health_check", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})
	v1.Post("/register", cc.Register)
	v1.Post("/emails", credential.NewMiddleware(cr), cc.SendEmail)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	app.Listen(fmt.Sprintf(":%d", port))
}
