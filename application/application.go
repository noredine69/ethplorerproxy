package application

import (
	"errors"
	"ethproxy/controllers"
	"ethproxy/services/config"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	ErrCannotStartApp = errors.New("Application cannot start: Config are broken")
)

type ApplicationInterface interface {
	Run()
	Stop()
}

type Application struct {
	configuration config.ConfigServiceInterface
	controller    *controllers.Controller
	sigs          chan os.Signal
}

func New(configFilePath string, debugMode bool) (*Application, error) {
	configuration, err := config.New(configFilePath, debugMode)
	if err != nil {
		log.Error().Msgf("Error loading configuration: %v", err)
		return nil, ErrCannotStartApp
	}

	app := &Application{
		configuration: configuration,
	}

	if err = app.initServiceLayer(); err != nil {
		log.Error().Msgf("Error intializing service layers: %v", err)
		return nil, ErrCannotStartApp
	}
	return app, nil
}

func (app *Application) Run() {
}

func (app *Application) Stop() {
}
func (app *Application) initServiceLayer() error {
	app.controller = controllers.New(app.configuration)
	app.controller.Run()
	return nil
}
