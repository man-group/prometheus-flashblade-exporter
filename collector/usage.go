// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"strconv"
)

type UsageCollector struct {
	fbClient *fb.FlashbladeClient
	fsFilterFlag string
	Usage    *prometheus.Desc
	Quota    *prometheus.Desc
}

func NewUsageCollector(fbClient *fb.FlashbladeClient, fsFilterFlag string) *UsageCollector {
	const subsystem = "usagequota"

	return &UsageCollector{
		fbClient: fbClient,
		fsFilterFlag: fsFilterFlag,
		Usage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "memory_usage_bytes"),
			"Usage of a user/group on a filesystem in bytes",
			[]string{"type", "name", "id", "filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
		Quota: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "memory_quota_bytes"),
			"Quota of a user/group on a filesystem in bytes",
			[]string{"type", "name", "id", "filesystem"},
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c UsageCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting Usage metrics", err)
	}
}

func (c UsageCollector) collect(ch chan<- prometheus.Metric) error {
	usage, err := c.fbClient.Usage(c.fsFilterFlag)
	if err != nil {
		return err
	}

	userUsage, groupUsage := usage.Users, usage.Groups

	for _, usage := range userUsage {
		for _, data := range usage.Items {
			ch <- prometheus.MustNewConstMetric(
				c.Usage,
				prometheus.GaugeValue,
				float64(data.Usage),
				"user", data.User.Name, strconv.FormatInt(int64(data.User.Id), 10), data.FileSystem["name"])

			ch <- prometheus.MustNewConstMetric(
				c.Quota,
				prometheus.GaugeValue,
				float64(data.Quota),
				"user", data.User.Name, strconv.FormatInt(int64(data.User.Id), 10), data.FileSystem["name"])
		}
	}

	for _, usage := range groupUsage {
		for _, data := range usage.Items {
			ch <- prometheus.MustNewConstMetric(
				c.Usage,
				prometheus.GaugeValue,
				float64(data.Usage),
				"group", data.Group.Name, strconv.FormatInt(int64(data.Group.Id), 10), data.FileSystem["name"])

			ch <- prometheus.MustNewConstMetric(
				c.Quota,
				prometheus.GaugeValue,
				float64(data.Quota),
				"group", data.Group.Name, strconv.FormatInt(int64(data.Group.Id), 10), data.FileSystem["name"])
		}
	}
	return nil
}
