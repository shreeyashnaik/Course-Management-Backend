package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/shreeyashnaik/Course-Management-Backend/common/views"
	"github.com/shreeyashnaik/Course-Management-Backend/config"
)

func CheckAuth(ctx *fiber.Ctx) error {
	authToken := ctx.Get("Authorization")
	authToken = strings.TrimSpace(strings.Replace(authToken, "Bearer ", "", 1))
	if authToken == "" {
		return views.UnAuthorisedView(ctx)
	}
	token, err := jwt.ParseWithClaims(authToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET_KEY), nil
	})
	if err != nil {
		return views.UnAuthorisedView(ctx)
	}

	claims := token.Claims.(jwt.MapClaims)

	userId, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return views.InternalServerError(ctx, err)
	}

	ctx.Locals("userId", userId)
	return ctx.Next()
}
