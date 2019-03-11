package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type AlertsCollector struct {
	fbClient   *fb.FlashbladeClient
	AlertTotal *prometheus.Desc
}

func NewAlertsCollector(fbClient *fb.FlashbladeClient) *AlertsCollector {
	const subsystem = "alert"

	return &AlertsCollector{
		fbClient: fbClient,
		AlertTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "alert_total"),
			"Number of open alerts",
			[]string{"severity"},
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c AlertsCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting alerts metrics", err)
	}
}

func countAlertsBySeverity(alerts fb.AlertsResponse) map[string]int {
	alertSeverityMap := make(map[string]int)

	for _, alert := range alerts.Items {
		alertSeverityMap[alert.Severity]++
	}

	return alertSeverityMap
}

func (c AlertsCollector) collect(ch chan<- prometheus.Metric) error {
	alerts, err := c.fbClient.OpenAlerts()
	if err != nil {
		return err
	}

	alertSeverityMap := countAlertsBySeverity(alerts)

	for sev, count := range alertSeverityMap {
		ch <- prometheus.MustNewConstMetric(
			c.AlertTotal,
			prometheus.GaugeValue,
			float64(count),
			sev)
	}

	return nil
}
