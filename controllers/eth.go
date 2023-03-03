package controllers

import (
	"ethproxy/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	METRICS_GET_LAST_BLOCK_END_POINT = "get_last_block_endpoint"
	METRICS_GET_LAST_BLOCK_NB_CALL   = "get_last_block_endpoint_nb_call"
	METRICS_GET_LAST_BLOCK_DESC      = "Metrics for Get Last Block endpoint"
)

func (ctrl *Controller) declareEthRoutes() {
	privateRoutes := ctrl.router.Group("/eth/")
	{
		privateRoutes.GET("lastblock", ctrl.lastBlock)
	}
	ctrl.addGaugeMetricForEndpoint(METRICS_GET_LAST_BLOCK_END_POINT, METRICS_GET_LAST_BLOCK_NB_CALL, METRICS_GET_LAST_BLOCK_DESC)
}

func (ctrl *Controller) lastBlock(c *gin.Context) {
	ctrl.incGaugeMetricForEndpoint(METRICS_GET_LAST_BLOCK_END_POINT, METRICS_GET_LAST_BLOCK_NB_CALL)
	getLastBlockResponse := models.LastBlock{}
	err := ctrl.ethApi.SendEthRequest(ctrl.ConfigService.GetConfig().Api.GetLastBlockFunction, []string{}, &getLastBlockResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The request has failed",
		})
	} else {
		c.JSON(http.StatusOK, getLastBlockResponse)
	}
}
