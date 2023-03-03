package models

type Service interface {
	GetType() ServiceType
}

type ServiceType string

const (
	BackEndAPI    ServiceType = "BackEndAPI"
	ConfigService ServiceType = "ConfigService"
)
