package config

import (
	"encoding/json"
	"ethproxy/models"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigServiceInterface interface {
	GetConfig() models.Config
	DebugMode() bool
}

type ConfigService struct {
	logger     zerolog.Logger
	config     models.Config
	configFile string
	debugMode  bool
}

func New(configFile string, debugMode bool) (*ConfigService, error) {
	service := ConfigService{
		configFile: configFile,
		debugMode:  debugMode,
	}
	service.logger = log.With().Str(models.LoggerService, string(service.GetType())).Logger()

	config, err := service.loadConfiguration()
	if err != nil {
		return nil, err
	}
	service.config = *config
	return &service, nil
}

func (service *ConfigService) GetType() models.ServiceType {
	return models.ConfigService
}

func (service *ConfigService) GetConfig() models.Config {
	return service.config
}

func (service *ConfigService) DebugMode() bool {
	return service.debugMode
}

// nolint: gosec
func (service *ConfigService) loadConfiguration() (ethConfig *models.Config, err error) {
	file, err := os.Open(service.configFile)
	if err != nil {
		service.logger.Error().Err(err).Msgf("Error opening %s", service.configFile)
		return
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()
	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&ethConfig)
	if err != nil {
		service.logger.Error().Err(err).Msgf("Error decoding content of config file")
	}
	return
}
