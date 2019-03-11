package collector

import (
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type ArrayPerformanceCollector struct {
	fbClient       *fb.FlashbladeClient
	BytesPerOp     *prometheus.Desc
	BytesPerRead   *prometheus.Desc
	BytesPerWrite  *prometheus.Desc
	OthersPerSec   *prometheus.Desc
	OutputPerSec   *prometheus.Desc
	ReadsPerSec    *prometheus.Desc
	UsecPerOtherOp *prometheus.Desc
	UsecPerReadOp  *prometheus.Desc
	UsecPerWriteOp *prometheus.Desc
	InputPerSec    *prometheus.Desc
	WritesPerSec   *prometheus.Desc
}

func NewArrayPerformanceCollector(fbClient *fb.FlashbladeClient) *ArrayPerformanceCollector {
	const subsystem = "perf"

	return &ArrayPerformanceCollector{
		fbClient: fbClient,
		BytesPerOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_op"),
			"Average operation size (read bytes+write bytes/(read ops+write ops))",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		BytesPerRead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_read"),
			"Average read size in bytes per read operation",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		BytesPerWrite: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "bytes_per_write"),
			"Average write size in bytes per write operation",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		OthersPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "others_per_sec"),
			"Other operations processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		OutputPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "output_per_sec"),
			"Bytes read per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reads_per_sec"),
			"Read requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerOtherOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_other_op"),
			"Average time, measured in microseconds, that the array takes to process other operations",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_op"),
			"Average time, measured in microseconds, that the array takes to process a read request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_op"),
			"Average time, measured in microseconds, that the array takes to process a write request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		InputPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "input_per_sec"),
			"Bytes written per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WritesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "writes_per_sec"),
			"Write requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c ArrayPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting array performance metrics", err)
	}
}

func (c ArrayPerformanceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.ArrayPerformance()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.BytesPerOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].BytesPerOp))

	ch <- prometheus.MustNewConstMetric(
		c.BytesPerRead,
		prometheus.GaugeValue,
		float64(stats.Items[0].BytesPerRead))

	ch <- prometheus.MustNewConstMetric(
		c.BytesPerWrite,
		prometheus.GaugeValue,
		float64(stats.Items[0].BytesPerWrite))

	ch <- prometheus.MustNewConstMetric(
		c.OthersPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].OthersPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.OutputPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].OutputPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReadsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerOtherOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerOtherOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteOp))

	ch <- prometheus.MustNewConstMetric(
		c.InputPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].InputPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WritesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WritesPerSec))

	return nil
}
