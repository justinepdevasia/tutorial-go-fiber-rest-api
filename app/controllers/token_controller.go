package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
)

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary     create a new access token
// @Tags        Token
// @Accept      json
// @Produce     json
// @Param       username body string true "Username"
// @Param       password body string true "Password"
// @Success     200 {string} status "ok"
// @Router      /api/v1/token/new [post]
func GetNewAccessToken(c *fiber.Ctx) error {
	// Get username and password from request body.
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
