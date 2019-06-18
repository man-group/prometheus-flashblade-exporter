// Copyright (c) 2019 Hudson River Trading LLC
// All rights reserved.

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

func NewFlashbladeCollector(fbClient *fb.FlashbladeClient) *FlashbladeCollector {
	alertsCollector := NewAlertsCollector(fbClient)
	arrayPerformanceCollector := NewArrayPerformanceCollector(fbClient)
	bladesCollector := NewBladesCollector(fbClient)
	filesystemsCollector := NewFilesystemsCollector(fbClient)
	usageCollector := NewUsageCollector(fbClient)
	fsPerformanceCollector := NewFSPerformanceCollector(fbClient)

	subcollectors := []Subcollector{
		alertsCollector,
		arrayPerformanceCollector,
		bladesCollector,
		filesystemsCollector,
		usageCollector,
		fsPerformanceCollector,
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
