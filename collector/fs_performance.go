// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/man-group/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type FSPerformanceCollector struct {
	fbClient         *fb.FlashbladeClient
	BytesPerOp       *prometheus.Desc
	BytesPerRead     *prometheus.Desc
	BytesPerWrite    *prometheus.Desc
	OthersPerSec     *prometheus.Desc
	ReadsPerSec      *prometheus.Desc
	ReadBytesPerSec  *prometheus.Desc
	UsecPerOtherOp   *prometheus.Desc
	UsecPerReadOp    *prometheus.Desc
	UsecPerWriteOp   *prometheus.Desc
	WritesPerSec     *prometheus.Desc
	WriteBytesPerSec *prometheus.Desc
}

func NewFSPerformanceCollector(fbClient *fb.FlashbladeClient) *FSPerformanceCollector {
	const subsystem = "fs_performance"

	return &FSPerformanceCollector{
		fbClient: fbClient,
		BytesPerOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_operation"),
			"Average operation size (read bytes+write bytes/(read operations + write operations))",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		BytesPerRead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_read"),
			"Average read size in bytes per read operation",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		BytesPerWrite: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_write"),
			"Average write size in bytes per write operation",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		OthersPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "non_read_write_operations_per_second"),
			"All operations except read processed per second",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reads_per_second"),
			"Read requests processed per second",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadBytesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "read_bytes_per_second"),
			"Read byte requests processed per second",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerOtherOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "seconds_per_non_read_write_operation"),
			"Average time, measured in seconds, that the array takes to process other operations",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "seconds_per_read_operation"),
			"Average time, measured in seconds, that the array takes to process a read request",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "seconds_per_write_operation"),
			"Average time, measured in seconds, that the array takes to process a write request",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),

		WritesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "writes_per_second"),
			"Write requests processed per second",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
		WriteBytesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_bytes_per_second"),
			"Write byte requests processed per second",
			[]string{"name"},
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c FSPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting quota metrics", err)
	}
}

func (c FSPerformanceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.FSPerformance()
	if err != nil {
		return err
	}

	for _, stat := range stats.Items {
		ch <- prometheus.MustNewConstMetric(
			c.BytesPerOp,
			prometheus.GaugeValue,
			float64(stat.BytesPerOp),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.BytesPerRead,
			prometheus.GaugeValue,
			float64(stat.BytesPerRead),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.BytesPerWrite,
			prometheus.GaugeValue,
			float64(stat.BytesPerWrite),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.OthersPerSec,
			prometheus.GaugeValue,
			float64(stat.OthersPerSec),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.ReadsPerSec,
			prometheus.GaugeValue,
			float64(stat.ReadsPerSec),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.ReadBytesPerSec,
			prometheus.GaugeValue,
			float64(stat.ReadBytesPerSec),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.UsecPerOtherOp,
			prometheus.GaugeValue,
			float64(stat.UsecPerOtherOp)/1e6,
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.UsecPerReadOp,
			prometheus.GaugeValue,
			float64(stat.UsecPerReadOp)/1e6,
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.UsecPerWriteOp,
			prometheus.GaugeValue,
			float64(stat.UsecPerWriteOp)/1e6,
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.WritesPerSec,
			prometheus.GaugeValue,
			float64(stat.WritesPerSec),
			stat.Name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.WriteBytesPerSec,
			prometheus.GaugeValue,
			float64(stat.WriteBytesPerSec),
			stat.Name,
		)
	}

	return nil
}
