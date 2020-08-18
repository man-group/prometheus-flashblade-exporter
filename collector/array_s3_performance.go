// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/man-group/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type ArrayS3PerformanceCollector struct {
	fbClient             *fb.FlashbladeClient
	ReadBucketsPerSec    *prometheus.Desc
	WriteBucketsPerSec   *prometheus.Desc
	ReadObjectsPerSec    *prometheus.Desc
	WriteObjectsPerSec   *prometheus.Desc
	OthersPerSec         *prometheus.Desc
	UsecPerReadBucketOp  *prometheus.Desc
	UsecPerWriteBucketOp *prometheus.Desc
	UsecPerReadObjectOp  *prometheus.Desc
	UsecPerWriteObjectOp *prometheus.Desc
	UsecPerOtherOp       *prometheus.Desc
}

func NewArrayS3PerformanceCollector(fbClient *fb.FlashbladeClient) *ArrayS3PerformanceCollector {
	const subsystem = "perf_s3"

	return &ArrayS3PerformanceCollector{
		fbClient: fbClient,
		ReadBucketsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "read_buckets_per_sec"),
			"Read bucket requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WriteBucketsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_buckets_per_sec"),
			"Write bucket requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadObjectsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "read_objects_per_sec"),
			"Read object requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WriteObjectsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "write_objects_per_sec"),
			"Write object requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		OthersPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "others_per_sec"),
			"Other operations processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadBucketOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_bucket_op"),
			"Average time, measured in microseconds, that the array takes to process a read bucket request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteBucketOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_bucket_op"),
			"Average time, measured in microseconds, that the array takes to process a write bucket request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadObjectOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_object_op"),
			"Average time, measured in microseconds, that the array takes to process a read object request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteObjectOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_object_op"),
			"Average time, measured in microseconds, that the array takes to process a write object request",
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

func (c ArrayS3PerformanceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.ArrayS3Performance()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.ReadBucketsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadBucketsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WriteBucketsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WriteBucketsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReadObjectsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadObjectsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WriteObjectsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WriteObjectsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.OthersPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].OthersPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadBucketOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadBucketOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteBucketOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteBucketOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadObjectOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadObjectOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteObjectOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteObjectOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerOtherOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerOtherOp))

	return nil
}

func (c ArrayS3PerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting array S3 performance metrics", err)
	}
}
