package controller

import (
    "errors"
    "github.com/ahmadirfaan/project-go/models/web"
    "github.com/ahmadirfaan/project-go/service"
    "github.com/ahmadirfaan/project-go/utils"
    "github.com/go-playground/validator"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type TransactionController interface {
    CreateTransaction(c *fiber.Ctx) error
    GetAllTransactionByUserId(c *fiber.Ctx) error
    UpdateTransaction(c *fiber.Ctx) error
    GiveAgentRating(c *fiber.Ctx) error
    DeleteTransactionById(c *fiber.Ctx) error
}

type transactionController struct {
    TransactionService service.TransactionService
}

func NewTransactionController(s service.TransactionService) TransactionController {
    return transactionController{
        TransactionService: s,
    }
}

func (ts transactionController) DeleteTransactionById(c *fiber.Ctx) error {
    transactionId := c.Params("transactionId")
    userIdToken, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    code, err := ts.TransactionService.DeleteTransaction(transactionId, userIdToken)
    if err != nil {
        if errors.As(err, &validator.ValidationErrors{}) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": utils.ValidatorErrors(err),
                "data":    nil,
            })
        } else {
            return c.Status(code).JSON(fiber.Map{
                "code":    code,
                "message": err.Error(),
                "data":    nil,
            })
        }
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code":    fiber.StatusOK,
        "message": "Success Delete Transaction",
        "data":    nil,
    })

}

func (ts transactionController) GiveAgentRating(c *fiber.Ctx) error {
    var request web.RequestRating
    transactionId := c.Params("transactionId")
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Error for handling your request",
            "data":    nil,
        })
    }
    userId, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    code, err := ts.TransactionService.GiveRatingTransaction(request, userId, transactionId)
    if err != nil {
        if errors.As(err, &validator.ValidationErrors{}) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": utils.ValidatorErrors(err),
                "data":    nil,
            })
        } else {
            return c.Status(code).JSON(fiber.Map{
                "code":    code,
                "message": err.Error(),
                "data":    nil,
            })
        }
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code":    fiber.StatusOK,
        "message": "Thank you for giving rating",
        "data":    nil,
    })
}

func (ts transactionController) GetAllTransactionByUserId(c *fiber.Ctx) error {
    userId, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    transactions, err := ts.TransactionService.GetAllTransactionByUserId(userId)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
    }
    if len(transactions) == 0 {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "code":    fiber.StatusOK,
            "message": "No Transaction Data",
            "data":    transactions,
        })
    } else {
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "code":    fiber.StatusOK,
            "message": nil,
            "data":    transactions,
        })
    }
}

func (ts transactionController) UpdateTransaction(c *fiber.Ctx) error {
    var request web.ChangeTransactionRequest
    transactionId := c.Params("transactionId")
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Error for handling your request",
            "data":    nil,
        })
    }
    userIdToken, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    userId, _ := strconv.Atoi(userIdToken)
    code, err := ts.TransactionService.ChangeStatusTransaction(transactionId, uint(userId), request)
    if err != nil {
        if errors.As(err, &validator.ValidationErrors{}) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code":    fiber.StatusBadRequest,
                "message": utils.ValidatorErrors(err),
                "data":    nil,
            })
        } else {
            return c.Status(code).JSON(fiber.Map{
                "code":    code,
                "message": err.Error(),
                "data":    nil,
            })
        }
    }
    var message string
    switch request.StatusTransaction {
    case 1:
        message = "Success Change to Agent on the way "
    case 2:
        message = "Success Change Status Transaction to Canceled"
    case 3:
        message = "Success Change to Done, Thank you for transaction"
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code":    fiber.StatusOK,
        "message": message,
        "data":    nil,
    })
}

func (ts transactionController) CreateTransaction(c *fiber.Ctx) error {
    var transaction web.CreateTransactionRequest
    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Error for handling your request",
            "data":    nil,
        })
    }
    userId, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    isAgent, err := ts.TransactionService.IsUserAgent(userId)
    if *isAgent {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "code":    fiber.StatusForbidden,
            "message": "You're not appropriate to do create transactions",
            "data":    nil,
        })
    }
    err = ts.TransactionService.CreateTransaction(transaction, userId)
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
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "code":    fiber.StatusCreated,
        "message": "Transaksi Diterima ",
        "data":    nil,
    })
}
