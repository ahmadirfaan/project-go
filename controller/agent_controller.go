package controller

import (
    "errors"
    "github.com/ahmadirfaan/project-go/models/web"
    "github.com/ahmadirfaan/project-go/service"
    "github.com/ahmadirfaan/project-go/utils"
    "github.com/go-playground/validator"
    "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
)

type AgentController interface {
    RegisterAgent(c *fiber.Ctx) error
    FindAgentByDistrictId(c *fiber.Ctx) error
}

type agentController struct {
    AgentService service.AgentService
}

func NewAgentController(s service.AgentService) AgentController {
    return agentController{
        AgentService: s,
    }
}

func (as agentController) RegisterAgent(c *fiber.Ctx) error {
    var agent web.RegisterAgentRequest
    if err := c.BodyParser(&agent); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Error for handling your request",
            "data":    nil,
        })
    }

    err := as.AgentService.RegisterAgent(agent)
    if err != nil {
        var mysqlErr *mysql.MySQLError
        if errors.As(err, &validator.ValidationErrors{}) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": utils.ValidatorErrors(err),
                "data":    nil,
            })
        } else if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "code":    fiber.StatusConflict,
                "message": "Username Already is exist",
                "data":    nil,
            })
        } else {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": err.Error(),
                "data":    nil,
            })
        }
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "code":    fiber.StatusCreated,
        "message": " Sukses Membuat Akun",
        "data":    nil,
    })
}

func (as agentController) FindAgentByDistrictId(c *fiber.Ctx) error {
    districtId := c.Query("districtId")
    users, err := as.AgentService.FindByDistrictId(districtId)
    if err != nil {
        if errors.As(err, &validator.ValidationErrors{}) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": utils.ValidatorErrors(err),
                "data":    nil,
            })
        } else {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": err.Error(),
                "data":    nil,
            })
        }
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code":    fiber.StatusOK,
        "message": nil,
        "data":    users,
    })
}
