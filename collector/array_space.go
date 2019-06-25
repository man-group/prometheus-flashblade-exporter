package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type ArraySpaceCollector struct {
	fbClient           *fb.FlashbladeClient
	Capacity           *prometheus.Desc
	VirtualUsage       *prometheus.Desc
	DataReduction      *prometheus.Desc
	UniqueUsage        *prometheus.Desc
	SnapshotUsage      *prometheus.Desc
	TotalPhysicalUsage *prometheus.Desc
}

func NewArraySpaceCollector(fbClient *fb.FlashbladeClient) *ArraySpaceCollector {
	const subsystem = "space"

	return &ArraySpaceCollector{
		fbClient: fbClient,
		Capacity: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "capacity_bytes"),
			"Usable capacity in bytes",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		VirtualUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "virtual_usage_bytes"),
			"Usage in bytes",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		DataReduction: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "data_reduction"),
			"Reduction of data",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UniqueUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "unique_usage_bytes"),
			"Physical usage in bytes",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		SnapshotUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "snapshot_usage_bytes"),
			"Physical usage by snapshots, non-unique",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		TotalPhysicalUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_usage_bytes"),
			"Total physical usage (including snapshots) in bytes",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c ArraySpaceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting array performance metrics", err)
	}
}

func (c ArraySpaceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.ArraySpace()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.Capacity,
		prometheus.GaugeValue,
		float64(stats.Items[0].Capacity))

	ch <- prometheus.MustNewConstMetric(
		c.VirtualUsage,
		prometheus.GaugeValue,
		float64(stats.Items[0].Space.Virtual))

	ch <- prometheus.MustNewConstMetric(
		c.DataReduction,
		prometheus.GaugeValue,
		float64(stats.Items[0].Space.DataReduction))

	ch <- prometheus.MustNewConstMetric(
		c.UniqueUsage,
		prometheus.GaugeValue,
		float64(stats.Items[0].Space.Unique))

	ch <- prometheus.MustNewConstMetric(
		c.SnapshotUsage,
		prometheus.GaugeValue,
		float64(stats.Items[0].Space.Snapshots))

	ch <- prometheus.MustNewConstMetric(
		c.TotalPhysicalUsage,
		prometheus.GaugeValue,
		float64(stats.Items[0].Space.TotalPhysical))

	return nil
}
