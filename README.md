# Pure FlashBlade Exporter for Prometheus

The exporter uses FlashBlade API version 1.2.

## Building

```
make build
```

Or simply:

```
make
```

## Usage

The exporter requires the name of the FlashBlade as a command-line argument:

`flashblade_exporter <flashblade>`

Pass in the token for API access as the environment variable `PUREFB_API`.  (An API token can be generated for
the FlashBlade by using the `pureadmin` command after SSHing to the device as 'pureuser'.)

The exporter accepts the following command line flags:

| Flag       | Description                                                              | Default |
| ---------- | ------------------------------------------------------------------------ | ------- |
| --port     | Port on which the exporter will bind to in order to serve up the metrics | 9130    |
| --insecure | Disable SSL verification                                                 | false   |


## Metrics

* Filesystem usage (unique, virtual, snapshot and total)
* Bandwidth, IOPS and latency for both read and write

```
# HELP flashblade_alert_num_critical_alerts Number of open critical severity alerts
# TYPE flashblade_alert_num_critical_alerts gauge
# HELP flashblade_alert_num_info_alerts Number of open info severity alerts
# TYPE flashblade_alert_num_info_alerts gauge
# HELP flashblade_alert_num_warning_alerts Number of open warning severity alerts
# TYPE flashblade_alert_num_warning_alerts gauge
# HELP flashblade_blade_num_healthy_blades Number of blades in healthy status
# TYPE flashblade_blade_num_healthy_blades gauge
# HELP flashblade_blade_num_unhealthy_blades Number of blades in a non-healthy status
# TYPE flashblade_blade_num_unhealthy_blades gauge
# HELP flashblade_collector_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which flashblade_collector was built.
# TYPE flashblade_collector_build_info gauge
# HELP flashblade_fs_data_reduction Reduction of data
# TYPE flashblade_fs_data_reduction gauge
# HELP flashblade_fs_snapshot_usage_bytes Physical usage by snapshots, non-unique
# TYPE flashblade_fs_snapshot_usage_bytes gauge
# HELP flashblade_fs_total_usage_bytes Total physical usage (including snapshots) in bytes
# TYPE flashblade_fs_total_usage_bytes gauge
# HELP flashblade_fs_unique_usage_bytes Physical usage in bytes
# TYPE flashblade_fs_unique_usage_bytes gauge
# HELP flashblade_fs_virtual_usage_bytes Usage in bytes
# TYPE flashblade_fs_virtual_usage_bytes gauge
# HELP flashblade_perf_bytes_per_op Average operation size (read bytes+write bytes/(read ops+write ops))
# TYPE flashblade_perf_bytes_per_op gauge
# HELP flashblade_perf_bytes_per_read Average read size in bytes per read operation
# TYPE flashblade_perf_bytes_per_read gauge
# HELP flashblade_perf_bytes_per_write Average write size in bytes per write operation
# TYPE flashblade_perf_bytes_per_write gauge
# HELP flashblade_perf_input_per_sec Bytes written per second
# TYPE flashblade_perf_input_per_sec gauge
# HELP flashblade_perf_others_per_sec Other operations processed per second
# TYPE flashblade_perf_others_per_sec gauge
# HELP flashblade_perf_output_per_sec Bytes read per second
# TYPE flashblade_perf_output_per_sec gauge
# HELP flashblade_perf_reads_per_sec Read requests processed per second
# TYPE flashblade_perf_reads_per_sec gauge
# HELP flashblade_perf_usec_per_other_op Average time, measured in microseconds, that the array takes to process other operations
# TYPE flashblade_perf_usec_per_other_op gauge
# HELP flashblade_perf_usec_per_read_op Average time, measured in microseconds, that the array takes to process a read request
# TYPE flashblade_perf_usec_per_read_op gauge
# HELP flashblade_perf_usec_per_write_op Average time, measured in microseconds, that the array takes to process a write request
# TYPE flashblade_perf_usec_per_write_op gauge
# HELP flashblade_perf_writes_per_sec Write requests processed per second
# TYPE flashblade_perf_writes_per_sec gauge
```

## TODO

- [ ] Exit the exporter on startup (instead of scrape) if the API token variable isn't set
- [ ] Add overall capacity metrics
- [ ] Add optional per-IP metrics
