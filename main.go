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
	insecureFlag   = kingpin.Flag("insecure", "Disable the verification of the SSL certificate").Default("false").Bool()
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
	kingpin.Version("0.1.0")
	kingpin.Parse()
	fbClient := fb.NewFlashbladeClient(*flashbladeFlag, *insecureFlag)
	fbCollector := collector.NewFlashbladeCollector(fbClient, *fsMetricFlag)
	prometheus.MustRegister(fbCollector)
	listen()
}
