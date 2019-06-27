// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type FilesystemsCollector struct {
	fbClient           *fb.FlashbladeClient
	VirtualUsage       *prometheus.Desc
	DataReduction      *prometheus.Desc
	UniqueUsage        *prometheus.Desc
	SnapshotUsage      *prometheus.Desc
	TotalPhysicalUsage *prometheus.Desc
}

func NewFilesystemsCollector(fbClient *fb.FlashbladeClient) *FilesystemsCollector {
	const subsystem = "fs"

	return &FilesystemsCollector{
		fbClient: fbClient,
		VirtualUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "virtual_usage_bytes"),
			"Usage in bytes",
			[]string{"filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
		DataReduction: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "data_reduction"),
			"Reduction of data",
			[]string{"filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
		UniqueUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "unique_usage_bytes"),
			"Physical usage in bytes",
			[]string{"filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
		SnapshotUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "snapshot_usage_bytes"),
			"Physical usage by snapshots, non-unique",
			[]string{"filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
		TotalPhysicalUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_usage_bytes"),
			"Total physical usage (including snapshots) in bytes",
			[]string{"filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c FilesystemsCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting blade metrics", err)
	}
}

func (c FilesystemsCollector) collect(ch chan<- prometheus.Metric) error {
	filesystems, err := c.fbClient.Filesystems()
	if err != nil {
		return err
	}

	for _, fs := range filesystems.Items {
		ch <- prometheus.MustNewConstMetric(
			c.VirtualUsage,
			prometheus.GaugeValue,
			float64(fs.Space.Virtual),
			fs.Name)

		ch <- prometheus.MustNewConstMetric(
			c.DataReduction,
			prometheus.GaugeValue,
			float64(fs.Space.DataReduction),
			fs.Name)

		ch <- prometheus.MustNewConstMetric(
			c.UniqueUsage,
			prometheus.GaugeValue,
			float64(fs.Space.Unique),
			fs.Name)

		ch <- prometheus.MustNewConstMetric(
			c.SnapshotUsage,
			prometheus.GaugeValue,
			float64(fs.Space.Snapshots),
			fs.Name)

		ch <- prometheus.MustNewConstMetric(
			c.TotalPhysicalUsage,
			prometheus.GaugeValue,
			float64(fs.Space.TotalPhysical),
			fs.Name)
	}

	return nil
}
