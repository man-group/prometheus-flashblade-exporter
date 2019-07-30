// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type S3BucketsCollector struct {
	fbClient           *fb.FlashbladeClient
	VirtualUsage       *prometheus.Desc
	DataReduction      *prometheus.Desc
	UniqueUsage        *prometheus.Desc
	SnapshotUsage      *prometheus.Desc
	TotalPhysicalUsage *prometheus.Desc
	ObjectCount        *prometheus.Desc
}

func NewS3BucketsCollector(fbClient *fb.FlashbladeClient) *S3BucketsCollector {
	const subsystem = "s3"

	return &S3BucketsCollector{
		fbClient: fbClient,
		VirtualUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "virtual_usage_bytes"),
			"Usage in bytes",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
		DataReduction: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "data_reduction"),
			"Reduction of data",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
		UniqueUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "unique_usage_bytes"),
			"Physical usage in bytes",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
		SnapshotUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "snapshot_usage_bytes"),
			"Physical usage by snapshots, non-unique",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
		TotalPhysicalUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_usage_bytes"),
			"Total physical usage (including snapshots) in bytes",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
		ObjectCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "object_count"),
			"The number of object within the bucket.",
			[]string{"bucket"},
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c S3BucketsCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting s3 blade metrics", err)
	}
}

func (c S3BucketsCollector) collect(ch chan<- prometheus.Metric) error {
	s3_buckets, err := c.fbClient.S3Buckets()
	if err != nil {
		return err
	}

	for _, s3_bucket := range s3_buckets.Items {
		ch <- prometheus.MustNewConstMetric(
			c.VirtualUsage,
			prometheus.GaugeValue,
			float64(s3_bucket.Space.Virtual),
			s3_bucket.Name)

		ch <- prometheus.MustNewConstMetric(
			c.DataReduction,
			prometheus.GaugeValue,
			float64(s3_bucket.Space.DataReduction),
			s3_bucket.Name)

		ch <- prometheus.MustNewConstMetric(
			c.UniqueUsage,
			prometheus.GaugeValue,
			float64(s3_bucket.Space.Unique),
			s3_bucket.Name)

		ch <- prometheus.MustNewConstMetric(
			c.SnapshotUsage,
			prometheus.GaugeValue,
			float64(s3_bucket.Space.Snapshots),
			s3_bucket.Name)

		ch <- prometheus.MustNewConstMetric(
			c.TotalPhysicalUsage,
			prometheus.GaugeValue,
			float64(s3_bucket.Space.TotalPhysical),
			s3_bucket.Name)

		ch <- prometheus.MustNewConstMetric(
			c.ObjectCount,
			prometheus.GaugeValue,
			float64(s3_bucket.ObjectCount),
			s3_bucket.Name)
	}

	return nil
}
