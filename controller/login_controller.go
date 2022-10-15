package controller

import (
	"errors"
	"github.com/ahmadirfaan/project-go/models/web"
	"github.com/ahmadirfaan/project-go/service"
	"github.com/ahmadirfaan/project-go/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type LoginController interface {
	Login(c *fiber.Ctx) error
	WelcomingAPI(c *fiber.Ctx) error
}

type loginController struct {
	LoginService service.LoginService
}

func NewLoginController(ls service.LoginService) LoginController {
	return loginController{
		LoginService: ls,
	}
}

func (cs loginController) Login(c *fiber.Ctx) error {
	var login web.LoginRequest
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	user, err := cs.LoginService.Login(login)
	if err != nil || user.Id == nil {
		if errors.As(err, &validator.ValidationErrors{}) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": utils.ValidatorErrors(err),
				"data":    nil,
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": err.Error(),
				"data":    nil,
			})
		}

	}

	token, expire, err := utils.GenerateToken(user)
	if err != nil { //Error that because of generates token
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Sucess Login",
		"data": fiber.Map{
			"accessToken": token,
			"expiredAt":   expire,
		},
	})
}

func (cs loginController) WelcomingAPI(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Welcome to API Agent-GO",
		"data":    nil,
	})
}
