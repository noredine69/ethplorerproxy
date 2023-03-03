package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	METRICS_HEALTH_DESC      = "Metrics for health (liveness, readyness) endpoint"
	METRICS_HEALTH_END_POINT = "healthz_endpoint"
	METRICS_HEALTH_NB_CALL   = "healthz_nb_call"
)

func (ctrl *Controller) declareHealthRoutes() {
	_ = ctrl.router.GET("/healthz", ctrl.healthz)
	ctrl.addGaugeMetricForEndpoint(METRICS_HEALTH_END_POINT, METRICS_HEALTH_NB_CALL, METRICS_HEALTH_DESC)
}

func (ctrl *Controller) healthz(c *gin.Context) {
	ctrl.incGaugeMetricForEndpoint(METRICS_HEALTH_END_POINT, METRICS_HEALTH_NB_CALL)
	c.Status(http.StatusOK)
}
