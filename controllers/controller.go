package controllers

import (
	"encoding/json"
	"ethproxy/services/config"
	"ethproxy/services/eth"
	"net/http"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	ConfigService  config.ConfigServiceInterface
	router         *gin.Engine
	ethApi         eth.EthAPIInterface
	metricsMonitor *ginmetrics.Monitor
}

func New(config config.ConfigServiceInterface) *Controller {
	ctrl := Controller{ConfigService: config}
	ctrl.ethApi = eth.New(config)
	ctrl.DeclareRoutes()
	return &ctrl
}

const (
	formatUintBase = 10
)

func (ctrl *Controller) DeclareRoutes() {
	ctrl.initGinEngine()
	ctrl.declareBackEndRoutes()
}

func (ctrl *Controller) declareBackEndRoutes() {
	// Don't change the order, metrics routes must be declare first in order to be called by other endpoints
	ctrl.declareMetricsRoutes()
	ctrl.declareEthRoutes()
	ctrl.declareHealthRoutes()
}

func (ctrl *Controller) initGinEngine() {
	if !ctrl.ConfigService.DebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}

	ctrl.router = gin.New()
	ctrl.router.Use(gin.Recovery())
	ctrl.router.Use(gin.LoggerWithFormatter(logWithZeroLog))

	// For profiling
	if ctrl.ConfigService.DebugMode() {
		pprof.Register(ctrl.router)
	}
}

func (ctrl *Controller) Run() {
	port := strconv.FormatUint(uint64(ctrl.ConfigService.GetConfig().Server.Port), formatUintBase)
	log.Info().Msg("Server Started on Port " + port)
	err := endless.ListenAndServe(":"+port,
		csrf.Protect([]byte(ctrl.ConfigService.GetConfig().Server.SecretKey),
			csrf.Secure(false),
			csrf.SameSite(csrf.SameSiteStrictMode),
			csrf.Path("/"),
			csrf.ErrorHandler(http.HandlerFunc(csrfErrorHandlerFunc)))(ctrl.router))

	if err != nil {
		log.Error().Err(err).Msgf("Error while starting the web server")
	}
}

func csrfErrorHandlerFunc(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusForbidden)
	msg := "CSRF Token is invalid"
	log.Error().Msg(msg)
	jsonStr, err := json.Marshal(msg)
	if err != nil {
		log.Error().Err(err).Msgf("Error while handling csrf token (Marshal)")
	}
	_, err = response.Write(jsonStr)
	if err != nil {
		log.Error().Err(err).Msgf("Error while handling csrf token (Write)")
	}
}
