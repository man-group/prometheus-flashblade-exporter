// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/man-group/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type BladesCollector struct {
	fbClient           *fb.FlashbladeClient
	NumHealthyBlades   *prometheus.Desc
	NumUnhealthyBlades *prometheus.Desc
}

func NewBladesCollector(fbClient *fb.FlashbladeClient) *BladesCollector {
	const subsystem = "blade"

	return &BladesCollector{
		fbClient: fbClient,
		NumHealthyBlades: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_healthy_blades"),
			"Number of blades in healthy status",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		NumUnhealthyBlades: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_unhealthy_blades"),
			"Number of blades in a non-healthy status",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c BladesCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting blade metrics", err)
	}
}

func countBladesByStatus(blades fb.BladesResponse) (int, int) {
	numHealthyBlades := 0
	numUnhealthyBlades := 0

	for _, blade := range blades.Items {
		switch blade.Status {
		case "healthy":
			numHealthyBlades++
		case "unused":
		default:
			numUnhealthyBlades++
		}
	}

	return numHealthyBlades, numUnhealthyBlades
}

func (c BladesCollector) collect(ch chan<- prometheus.Metric) error {
	blades, err := c.fbClient.Blades()
	if err != nil {
		return err
	}

	numHealthyBlades, numUnhealthyBlades := countBladesByStatus(blades)

	ch <- prometheus.MustNewConstMetric(
		c.NumHealthyBlades,
		prometheus.GaugeValue,
		float64(numHealthyBlades))

	ch <- prometheus.MustNewConstMetric(
		c.NumUnhealthyBlades,
		prometheus.GaugeValue,
		float64(numUnhealthyBlades))

	return nil
}
