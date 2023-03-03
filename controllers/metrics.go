package controllers

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

const (
	slowTime = 10
)

func (ctrl *Controller) declareMetricsRoutes() {
	ctrl.metricsMonitor = ginmetrics.GetMonitor()
	ctrl.metricsMonitor.SetMetricPath("/metrics")
	ctrl.metricsMonitor.SetSlowTime(slowTime)
	ctrl.metricsMonitor.Use(ctrl.router)
}

func (ctrl *Controller) addGaugeMetricForEndpoint(metricName, metricFieldName, metricDescription string) {
	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Gauge,
		Name:        metricName,
		Description: metricDescription,
		Labels:      []string{metricFieldName},
	}
	_ = ctrl.metricsMonitor.AddMetric(gaugeMetric)
}

func (ctrl *Controller) incGaugeMetricForEndpoint(metricName, metricFieldName string) {
	_ = ctrl.metricsMonitor.GetMetric(metricName).Inc([]string{metricFieldName})
}
