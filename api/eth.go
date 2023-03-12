package api

import (
	"ethproxy/services/eth"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	METRICS_GET_LAST_BLOCK_END_POINT = "get_last_block_endpoint"
	METRICS_GET_LAST_BLOCK_NB_CALL   = "get_last_block_endpoint_nb_call"
	METRICS_GET_LAST_BLOCK_DESC      = "Metrics for Get Last Block endpoint"
)

func (api *Api) declareEthRoutes() {
	privateRoutes := api.router.Group("/eth/")
	{
		privateRoutes.GET("lastblock", api.lastBlock)
	}
	api.addGaugeMetricForEndpoint(METRICS_GET_LAST_BLOCK_END_POINT, METRICS_GET_LAST_BLOCK_NB_CALL, METRICS_GET_LAST_BLOCK_DESC)
}

func (api *Api) lastBlock(ginContext *gin.Context) {
	api.incGaugeMetricForEndpoint(METRICS_GET_LAST_BLOCK_END_POINT, METRICS_GET_LAST_BLOCK_NB_CALL)
	getLastBlockResponse := eth.LastBlock{}
	err := api.ethApi.Request(ginContext, api.config.Eth.GetLastBlockFunction, []string{}, &getLastBlockResponse)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"message": "The request has failed",
		})
	} else {
		ginContext.JSON(http.StatusOK, getLastBlockResponse)
	}
}
