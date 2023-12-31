package api

import (
	database "disco/db"
	"disco/structs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TokenMiddleware(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()
		token := headers["Token"]

		if len(token) <= 0 {
			ctx.Status(403)
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid Token, Unauthenticated",
			})
		}

		if user, err := database.GetUserByToken(db, token); err != nil {
			ctx.Status(500)
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Internal Server Error",
			})
		} else {
			if (user != structs.User{}) {
				ctx.Locals("User", user)
			} else {
				ctx.Status(403)
				return ctx.JSON(structs.Response{
					Success: false,
					Message: "Unauthenticated",
				})
			}
		}

		return ctx.Next()
	}
}
