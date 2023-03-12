package config

type Config struct {
	Api       ApiConfig
	Eth       EthConfig
	DebugMode bool
}

type EthConfig struct {
	Url                  string `json:"url"`
	ApiKey               string `json:"apiKey"`
	GetLastBlockFunction string `json:"function"`
}

type ApiConfig struct {
	Port      int    `json:"port"`
	SecretKey string `json:"secretKey"`
}
