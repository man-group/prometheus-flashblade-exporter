//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
//
// Copyright (c) 2019 Hudson River Trading LLC
// All rights reserved.
//

package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"strconv"
)

type UsageCollector struct {
	fbClient *fb.FlashbladeClient
	Usage    *prometheus.Desc
	Quota    *prometheus.Desc
}

func NewUsageCollector(fbClient *fb.FlashbladeClient) *UsageCollector {
	const subsystem = "usagequota"

	return &UsageCollector{
		fbClient: fbClient,
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
	usage, err := c.fbClient.Usage()
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
