package config

import (
	"github.com/rs/zerolog/log"
)

const NoMockError = "No mocked function provided %s"

type ConfigServiceMock struct {
	GetConfigFunc func() Config
	DebugModeFunc func() bool
}

func NewMock() *ConfigServiceMock {
	return &ConfigServiceMock{}
}

func (mock *ConfigServiceMock) GetConfig() Config {
	if mock.GetConfigFunc != nil {
		return mock.GetConfigFunc()
	}
	log.Warn().Msgf(NoMockError, "GetConfig")
	return Config{}
}
func (mock *ConfigServiceMock) DebugMode() bool {
	if mock.DebugModeFunc != nil {
		return mock.DebugModeFunc()
	}
	log.Warn().Msgf(NoMockError, "DebugMode")
	return false
}
