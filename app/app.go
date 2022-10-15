package app

import "github.com/ahmadirfaan/project-go/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
