// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"sync"

	"github.com/manahl/prometheus-flashblade-exporter/fb"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Used by the Prometheus library
	namespace = "flashblade"
)

type Subcollector interface {
	Collect(ch chan<- prometheus.Metric)
}

type FlashbladeCollector struct {
	subcollectors []Subcollector
}

func NewFlashbladeCollector(fbClient *fb.FlashbladeClient, fsMetricFlag bool) *FlashbladeCollector {
	alertsCollector := NewAlertsCollector(fbClient)
	arrayPerformanceCollector := NewArrayPerformanceCollector(fbClient)
	arraySpaceCollector := NewArraySpaceCollector(fbClient)
	bladesCollector := NewBladesCollector(fbClient)
	filesystemsCollector := NewFilesystemsCollector(fbClient)
	s3BucketsCollector := NewS3BucketsCollector(fbClient)

	subcollectors := []Subcollector{
		alertsCollector,
		arrayPerformanceCollector,
		arraySpaceCollector,
		bladesCollector,
		filesystemsCollector,
		s3BucketsCollector,
	}

	if fsMetricFlag {
		usageCollector := NewUsageCollector(fbClient)
		fsPerformanceCollector := NewFSPerformanceCollector(fbClient)

		subcollectors = append(subcollectors, usageCollector, fsPerformanceCollector)
	}

	return &FlashbladeCollector{subcollectors: subcollectors}
}

func (fbCollector FlashbladeCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(fbCollector.subcollectors))
	for _, c := range fbCollector.subcollectors {
		go func(c Subcollector) {
			c.Collect(ch)
			wg.Done()
		}(c)
	}
	wg.Wait()
}

func (fbCollector FlashbladeCollector) Describe(ch chan<- *prometheus.Desc) {

}
