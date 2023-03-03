package models

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
