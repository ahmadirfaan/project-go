package cli

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ahmadirfaan/project-go/middleware"
	route "github.com/ahmadirfaan/project-go/routes"

	"github.com/ahmadirfaan/project-go/app"
	"github.com/ahmadirfaan/project-go/config"
	databaseconn "github.com/ahmadirfaan/project-go/config/database"
	"github.com/ahmadirfaan/project-go/controller"
	"github.com/ahmadirfaan/project-go/repositories"
	"github.com/ahmadirfaan/project-go/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(application *app.Application) {
	fiberConfig := config.FiberConfig()
	appFiber := fiber.New(fiberConfig)

	// set up connection
	db := databaseconn.InitDb()
	//Repository
	userRepo := repositories.NewUserRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	provinceRepo := repositories.NewProvinceRepository(db)
	regencyRepo := repositories.NewRegencyRepository(db)
	agentRepo := repositories.NewAgentRepository(db)
	districtRepo := repositories.NewDistrictRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionTypeRepo := repositories.NewTransactionTypeRepository(db)

	// Service
	customerService := service.NewCustomerService(customerRepo, userRepo, db)
	locationService := service.NewLocationService(provinceRepo, regencyRepo, districtRepo)
	loginService := service.NewLoginService(userRepo)
	agentService := service.NewAgentService(agentRepo, userRepo, locationService, db)
	transactionService := service.NewTransactionService(
		transactionRepo, transactionTypeRepo, locationService, userRepo, agentRepo, db)

	// Controller
	customerController := controller.NewCustomerController(customerService)
	locationController := controller.NewLocationController(locationService)
	loginController := controller.NewLoginController(loginService)
	agentController := controller.NewAgentController(agentService)
	transactionController := controller.NewTransactionController(transactionService)

	// Route
	middleware.LoggerRoute(appFiber)
	middleware.AllowCrossOrigin(appFiber)
	appFiber.Post("/login", loginController.Login)
	appFiber.Post("/customers", customerController.RegisterCustomer)
	appFiber.Post("/agents", agentController.RegisterAgent)
	appFiber.Get("/", loginController.WelcomingAPI)
	appFiber.Get("/locations/districts", locationController.GetAllDistrictsByRegencyId)
	appFiber.Get("/locations/provinces", locationController.GetAllProvinces)
	appFiber.Get("/locations", locationController.GetAllRegenciesByProvinceId)
	// Set Middleware Auth with JWT Config
	middleware.MiddlewareAuth(appFiber)
	appFiber.Post("/transactions", transactionController.CreateTransaction)
	appFiber.Get("/transactions", transactionController.GetAllTransactionByUserId)
	appFiber.Put("/transactions/:transactionId", transactionController.UpdateTransaction)
	appFiber.Put("/transactions/rating/:transactionId", transactionController.GiveAgentRating)
	appFiber.Get("/agents", agentController.FindAgentByDistrictId)
	appFiber.Delete("/transactions/:transactionId", transactionController.DeleteTransactionById)
	route.NotFoundRoute(appFiber)
	StartServerWithGracefulShutdown(appFiber, application.Config.AppPort)
}

func StartServerWithGracefulShutdown(app *fiber.App, port string) {
	appPort := fmt.Sprintf(`:%s`, port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := app.Listen(appPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
