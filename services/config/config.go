package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigServiceInterface interface {
	GetConfig() Config
	DebugMode() bool
}

type ConfigService struct {
	logger     zerolog.Logger
	config     Config
	configFile string
	debugMode  bool
}

type Config struct {
	Api       ApiConfig
	Server    HttServerConfig
	DebugMode bool
}

type ApiConfig struct {
	Url                  string `json:"url"`
	ApiKey               string `json:"apiKey"`
	GetLastBlockFunction string `json:"function"`
}

type HttServerConfig struct {
	Port      int    `json:"port"`
	SecretKey string `json:"secretKey"`
}

func New(configFile string, debugMode bool) (*ConfigService, error) {
	service := ConfigService{
		configFile: configFile,
		debugMode:  debugMode,
	}

	config, err := service.loadConfiguration()
	if err != nil {
		return nil, err
	}
	service.config = *config
	return &service, nil
}

func (service *ConfigService) GetConfig() Config {
	return service.config
}

func (service *ConfigService) DebugMode() bool {
	return service.debugMode
}

// nolint: gosec
func (service *ConfigService) loadConfiguration() (ethConfig *Config, err error) {
	file, err := os.Open(service.configFile)
	if err != nil {
		log.Error().Err(err).Msgf("Error opening %s", service.configFile)
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
		log.Error().Err(err).Msgf("Error decoding content of config file")
	}
	return
}
