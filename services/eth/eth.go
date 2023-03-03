package eth

import (
	"encoding/json"
	"ethproxy/services/backendapi"
	"ethproxy/services/config"
	"fmt"

	"github.com/rs/zerolog/log"
)

type EthAPIInterface interface {
	SendEthRequest(methodName string, params []string, response interface{}) error
}

type EthAPI struct {
	Config     config.ConfigServiceInterface
	backendapi backendapi.BackEndAPIInterface
}

func New(configService config.ConfigServiceInterface) *EthAPI {
	return &EthAPI{
		Config:     configService,
		backendapi: backendapi.New(configService),
	}
}

func (eth *EthAPI) SendEthRequest(methodName string, params []string, response interface{}) error {
	requestUrl := fmt.Sprintf("%s%s?apiKey=%s", eth.Config.GetConfig().Api.Url, eth.Config.GetConfig().Api.GetLastBlockFunction, eth.Config.GetConfig().Api.ApiKey)
	resp, err := eth.backendapi.BuildAndSendGetRequest(requestUrl)
	defer eth.backendapi.CloseRespBody(requestUrl, resp)
	if err != nil {
		log.Error().Err(err).Msgf("Error from buildAndSendRequest")
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("Error while decoding eth response json string")
		return err
	}
	log.Info().Msgf("Received Eth response : %v", response)

	return nil
}
