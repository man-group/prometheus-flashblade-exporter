// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/man-group/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type ArrayHttpPerformanceCollector struct {
	fbClient           *fb.FlashbladeClient
	ReadDirsPerSec     *prometheus.Desc
	WriteDirsPerSec    *prometheus.Desc
	ReadFilesPerSec    *prometheus.Desc
	WriteFilesPerSec   *prometheus.Desc
	OthersPerSec       *prometheus.Desc
	UsecPerReadDirOp   *prometheus.Desc
	UsecPerWriteDirOp  *prometheus.Desc
	UsecPerReadFileOp  *prometheus.Desc
	UsecPerWriteFileOp *prometheus.Desc
	UsecPerOtherOp     *prometheus.Desc
}

func NewArrayHttpPerformanceCollector(fbClient *fb.FlashbladeClient) *ArrayHttpPerformanceCollector {
	const subsystem = "perf_http"

	return &ArrayHttpPerformanceCollector{
		fbClient: fbClient,
		ReadDirsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "read_dirs_per_sec"),
			"Read directories requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WriteDirsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_dirs_per_sec"),
			"Write directories requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadFilesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "read_files_per_sec"),
			"Read files requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WriteFilesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_files_per_sec"),
			"Write files requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		OthersPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "others_per_sec"),
			"Other operations processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadDirOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_dir_op"),
			"Average time, measured in microseconds, that the array takes to process a read directory request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteDirOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_dir_op"),
			"Average time, measured in microseconds, that the array takes to process a write directory request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadFileOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_file_op"),
			"Average time, measured in microseconds, that the array takes to process a read file request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteFileOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_file_op"),
			"Average time, measured in microseconds, that the array takes to process a write file request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerOtherOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_other_op"),
			"Average time, measured in microseconds, that the array takes to process other operations",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}
func (c ArrayHttpPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting array HTTP performance metrics", err)
	}
}

func (c ArrayHttpPerformanceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.ArrayHttpPerformance()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.ReadDirsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadDirsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WriteDirsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WriteDirsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReadFilesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadFilesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WriteFilesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WriteFilesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.OthersPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].OthersPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadDirOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadDirOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteDirOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteDirOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadFileOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadFileOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteFileOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteFileOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerOtherOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerOtherOp))

	return nil
}
