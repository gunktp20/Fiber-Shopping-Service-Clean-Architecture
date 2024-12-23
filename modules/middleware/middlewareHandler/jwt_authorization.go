package handler

import (
	"net/http"
	"shopping-service-be/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *middlewareHandler) JwtAuthorization(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return response.ErrResponse(c, http.StatusUnauthorized, "Authorization header is missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return response.ErrResponse(c, http.StatusUnauthorized, "Invalid authorization format")
	}

	_, err := h.middlewareUsecase.JwtAuthorization(tokenParts[1])
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return c.Next()

}
