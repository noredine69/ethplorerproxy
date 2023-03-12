package eth

import (
	"context"
	"encoding/json"
	"ethproxy/services/config"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type LastBlock struct {
	LastBlockId uint `json:"lastBlock"`
}

type ProxyEthRequest struct {
	JsonRpcVersion string   `json:"jsonrpc"`
	Method         string   `json:"method"`
	Params         []string `json:"params"`
	Id             uint     `json:"id"`
}

type ProxyEthResponse struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Id             uint   `json:"id"`
	Result         string `json:"result"`
}

type EthAPIInterface interface {
	Request(ctx context.Context, methodName string, params []string, response interface{}) error
}

type EthAPI struct {
	config config.EthConfig
}

func New(config config.EthConfig) *EthAPI {
	return &EthAPI{
		config: config,
	}
}

func (eth *EthAPI) Request(ctx context.Context, methodName string, params []string, response interface{}) error {
	requestUrl := fmt.Sprintf("%s%s?apiKey=%s", eth.config.Url, eth.config.GetLastBlockFunction, eth.config.ApiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("received bad http status for %s, %d", requestUrl, resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("error while decoding eth response, %w", err)
	}
	log.Info().Msgf("received eth response : %+v", response)
	return nil
}
