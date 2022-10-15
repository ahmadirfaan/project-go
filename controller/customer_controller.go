package controller

import (
    "errors"
    "github.com/ahmadirfaan/project-go/models/web"
    "github.com/ahmadirfaan/project-go/utils"
    "github.com/go-playground/validator"
    "github.com/go-sql-driver/mysql"
    "log"

    "github.com/ahmadirfaan/project-go/service"
    "github.com/gofiber/fiber/v2"
)

type CustomerController interface {
    RegisterCustomer(c *fiber.Ctx) error
}

type customerController struct {
    CustomerService service.CustomerService
}

//NewCustomerController -> returns new customer controller
func NewCustomerController(s service.CustomerService) CustomerController {
    return customerController{
        CustomerService: s,
    }
}

func (cs customerController) RegisterCustomer(c *fiber.Ctx) error {
    log.Print("[CustomerController]...add Customer")
    var customer web.RegisterCustomerRequest
    if err := c.BodyParser(&customer); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
    }

    err := cs.CustomerService.RegisterCustomer(customer)
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
        "message": "Sukses Membuat Akun",
        "data":    nil,
    })
}
