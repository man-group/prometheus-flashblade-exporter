# Pure FlashBlade Exporter for Prometheus

Export Prometheus scrapable metrics from a Pure Storage FlashBlade. 

The exporter minimally requires FlashBlade API version 1.2 (and version 1.3 for S3 metrics).

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

Pass in the token for API access as the environment variable `PUREFB_API`. (An API token can be generated for
the FlashBlade by using the `pureadmin` command after SSHing to the device as 'pureuser'.)

The exporter accepts the following command line flags:

| Flag                 | Description                                                                                    | Default |
| -------------------- | ---------------------------------------------------------------------------------------------- | ------- |
| --port               | Port on which the exporter will bind to in order to serve up the metrics                       | 9130    |
| --insecure           | Disable SSL verification                                                                       | false   |
| --filesystem-metrics | Enable per-filesystem performance and user/group metrics (requires FlashBlade API version 1.8) | false   |


## Metrics

* Filesystem usage (unique, virtual, snapshot and total)
* S3 bucket usage (unique, virtual, snapshot, total and number of objects)
* Bandwidth, IOPS and latency for both read and write
* Number of alerts of each severity
* FlashBlade total capacity and usage
* Usage statistics for each user and group per filesystem (with `--filesystem-metrics`)
* Filesystem performance (with `--filesystem-metrics`)

```
# HELP flashblade_alert_num_open Number of open alerts of each severity
# TYPE flashblade_alert_num_open gauge
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
# HELP flashblade_s3_data_reduction Reduction of data
# TYPE flashblade_s3_data_reduction gauge
# HELP flashblade_s3_object_count The number of object within the bucket.
# TYPE flashblade_s3_object_count gauge
# HELP flashblade_s3_snapshot_usage_bytes Physical usage by snapshots, non-unique
# TYPE flashblade_s3_snapshot_usage_bytes gauge
# HELP flashblade_s3_total_usage_bytes Total physical usage (including snapshots) in bytes
# TYPE flashblade_s3_total_usage_bytes gauge
# HELP flashblade_s3_unique_usage_bytes Physical usage in bytes
# TYPE flashblade_s3_unique_usage_bytes gauge
# HELP flashblade_s3_virtual_usage_bytes Usage in bytes
# TYPE flashblade_s3_virtual_usage_bytes gauge
# HELP flashblade_space_capacity_bytes Usable capacity in bytes
# TYPE flashblade_space_capacity_bytes gauge
# HELP flashblade_space_data_reduction Reduction of data
# TYPE flashblade_space_data_reduction gauge
# HELP flashblade_space_snapshot_usage_bytes Physical usage by snapshots, non-unique
# TYPE flashblade_space_snapshot_usage_bytes gauge
# HELP flashblade_space_total_usage_bytes Total physical usage (including snapshots) in bytes
# TYPE flashblade_space_total_usage_bytes gauge
# HELP flashblade_space_unique_usage_bytes Physical usage in bytes
# TYPE flashblade_space_unique_usage_bytes gauge
# HELP flashblade_space_virtual_usage_bytes Usage in bytes
# TYPE flashblade_space_virtual_usage_bytes gauge
```

Additionally, with `--filesystem-metrics`:

```
# HELP flashblade_usagequota_memory_quota_bytes Quota of a user/group on a filesystem in bytes
# TYPE flashblade_usagequota_memory_quota_bytes gauge
# HELP flashblade_usagequota_memory_usage_bytes Usage of a user/group on a filesystem in bytes
# TYPE flashblade_usagequota_memory_usage_bytes gauge
# HELP flashblade_fs_performance_bytes_per_op Average operation size (read bytes+write bytes/(read ops+write ops))
# TYPE flashblade_fs_performance_bytes_per_op gauge
# HELP flashblade_fs_performance_bytes_per_read Average read size in bytes per read operation
# TYPE flashblade_fs_performance_bytes_per_read gauge
# HELP flashblade_fs_performance_bytes_per_write Average write size in bytes per write operation
# TYPE flashblade_fs_performance_bytes_per_write gauge
# HELP flashblade_fs_performance_non_read_write_operations_per_second All operations except read processed per second
# TYPE flashblade_fs_performance_non_read_write_operations_per_second gauge
# HELP flashblade_fs_performance_read_bytes_per_second Read byte requests processed per second
# TYPE flashblade_fs_performance_read_bytes_per_second gauge
# HELP flashblade_fs_performance_reads_per_second Read requests processed per second
# TYPE flashblade_fs_performance_reads_per_second gauge
# HELP flashblade_fs_performance_sec_per_non_read_write_op Average time, measured in seconds, that the array takes to process other operations
# TYPE flashblade_fs_performance_sec_per_non_read_write_op gauge
# HELP flashblade_fs_performance_sec_per_read_op Average time, measured in seconds, that the array takes to process a read request
# TYPE flashblade_fs_performance_sec_per_read_op gauge
# HELP flashblade_fs_performance_sec_per_write_op Average time, measured in seconds, that the array takes to process a write request
# TYPE flashblade_fs_performance_sec_per_write_op gauge
# HELP flashblade_fs_performance_write_bytes_per_second Write byte requests processed per second
# TYPE flashblade_fs_performance_write_bytes_per_second gauge
# HELP flashblade_fs_performance_writes_per_second Write requests processed per second
# TYPE flashblade_fs_performance_writes_per_second gauge
```

## Authors
Prometheus FlashBlade Exporter has been under development since 2019 and welcomes contributions.

* [Michael Captain](https://github.com/macaptain), Man Group
* [Advait Bhatwdekar](https://github.com/You-NeverKnow), Hudson River Trading LLC
* [Jeff Patti](https://github.com/jepatti), Man Group
