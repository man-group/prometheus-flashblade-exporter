// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/manahl/prometheus-flashblade-exporter/collector"
	"github.com/manahl/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	flashbladeFlag = kingpin.Arg("flashblade", "Address of the target Flashblade.").Required().String()
	portFlag       = kingpin.Flag("port", "Port to listen on.").Short('p').Default("9130").String()
	fsMetricFlag   = kingpin.Flag("filesystem-metrics", "Export filesystem and usage data metrics for each user and group.").Default("false").Bool()
	fsFilterFlag   = kingpin.Flag("filesystem-filter-regexp", "Regexp limiting the filesystems for which metrics are exported").Default(".*").String()
	insecureFlag   = kingpin.Flag("insecure", "Disable the verification of the SSL certificate").Default("false").Bool()
	apiVersionFlag = kingpin.Flag("api-version", "API version to query the flashblade").Default("1.7").String()
)

func init() {
	prometheus.MustRegister(version.NewCollector("flashblade_collector"))
}

func listen() {
	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
	})
	log.Printf("Starting metrics gathering for FlashBlade %v on port %v", *flashbladeFlag, *portFlag)
	log.Fatal(http.ListenAndServe(":"+string(*portFlag), nil))
}

func main() {
	kingpin.Version("0.3.0")
	kingpin.Parse()
	fbClient := fb.NewFlashbladeClient(*flashbladeFlag, *insecureFlag, *apiVersionFlag)
	fbCollector := collector.NewFlashbladeCollector(fbClient, *fsMetricFlag, *fsFilterFlag)
	prometheus.MustRegister(fbCollector)
	listen()
}
