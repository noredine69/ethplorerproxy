package controllers

import (
	"ethproxy/models"
	"ethproxy/services/config"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	conf := initConfigHelperForHealth()
	controller := New(conf)
	ts := httptest.NewServer(controller.router)
	defer ts.Close()
	body := checkLogsRouteCallStatusOk(t, fmt.Sprintf("%s/healthz", ts.URL), "GET")
	assert.Equal(t, "", body)
}

func initConfigHelperForHealth() config.ConfigServiceInterface {
	conf := config.NewMock()
	conf.GetConfigFunc = func() models.Config {
		return models.Config{}
	}

	return conf
}
