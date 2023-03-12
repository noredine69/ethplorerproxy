package backendapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"ethproxy/services/config"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	errorFromEthServer = errors.New("impossible to get a response from eth server")
)

type BackEndAPIInterface interface {
	BuildAndSendPostRequest(urlRequest string, request interface{}) (resp *http.Response, err error)
	BuildAndSendGetRequest(urlRequest string) (resp *http.Response, err error)
	CloseRespBody(url string, resp *http.Response)
}

type BackEndAPI struct {
	Config     config.ConfigServiceInterface
	HTTPClient http.Client
}

const (
	httpTimeOut = 1500 * time.Millisecond
)

var (
	ErrNotOk error = errors.New("HTTP Response is not 200")
)

func New(configService config.ConfigServiceInterface) *BackEndAPI {
	return &BackEndAPI{
		Config:     configService,
		HTTPClient: http.Client{},
	}
}

func (back *BackEndAPI) BuildAndSendGetRequest(urlRequest string) (resp *http.Response, err error) {
	log.Error().Msgf("BuildAndSendRequest %s", urlRequest)
	ctx, cancelCtx := context.WithTimeout(context.Background(), httpTimeOut)
	defer cancelCtx()
	req, errReq := http.NewRequestWithContext(ctx, "GET", urlRequest, nil)
	if errReq != nil {
		log.Error().Err(errReq).Msgf("Error while creating a new request")
		return nil, errReq
	}
	ctx.Done()

	//nolint: bodyclose
	resp, err = back.HTTPClient.Do(req)

	if err != nil {
		log.Error().Err(err).Msgf("Error while trying to contact eth server")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Got HTTP response code: %v", resp.StatusCode)
		log.Error().Err(errorFromEthServer).Msg(msg)
		return nil, errorFromEthServer
	}
	return resp, nil
}

func (back *BackEndAPI) BuildAndSendPostRequest(urlRequest string, request interface{}) (resp *http.Response, err error) {
	var jsonStr []byte
	jsonStr, errMarshall := json.Marshal(request)
	if errMarshall != nil {
		log.Error().Err(errMarshall).Msgf("Error while marshaling request : %v", request)
		return nil, errMarshall
	}
	log.Error().Msgf("BuildAndSendRequest %s", urlRequest)
	ctx, cancelCtx := context.WithTimeout(context.Background(), httpTimeOut)
	defer cancelCtx()
	req, errReq := http.NewRequestWithContext(ctx, "POST", urlRequest, bytes.NewBuffer(jsonStr))
	if errReq != nil {
		log.Error().Err(errReq).Msgf("Error while creating a new request")
		return nil, errReq
	}
	req.Header.Set("Content-Type", "application/json")

	//nolint: bodyclose
	resp, err = back.HTTPClient.Do(req)

	if err != nil {
		log.Error().Err(err).Msgf("Error while trying to contact eth server")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Got HTTP response code: %v", resp.StatusCode)
		log.Error().Err(errorFromEthServer).Msg(msg)
		return nil, errorFromEthServer
	}
	return resp, nil
}

func (back *BackEndAPI) CloseRespBody(url string, resp *http.Response) {
	if resp == nil {
		return
	}
	_, errDiscard := io.Copy(ioutil.Discard, resp.Body)
	if errDiscard != nil {
		log.Warn().Err(errDiscard).Msgf("Discard body %s failed", url)
	}
	errClose := resp.Body.Close()
	if errClose != nil {
		log.Error().Err(errClose).Msgf("Closing body %s failed", url)
	}
}
