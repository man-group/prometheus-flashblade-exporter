// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package collector

import (
	"github.com/man-group/prometheus-flashblade-exporter/fb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type ArrayNfsPerformanceCollector struct {
	fbClient                             *fb.FlashbladeClient
	AccessesPerSec                       *prometheus.Desc
	AggregateFileMetadataCreatesPerSec   *prometheus.Desc
	AggregateFileMetadataModifiesPerSec  *prometheus.Desc
	AggregateFileMetadataReadsPerSec     *prometheus.Desc
	AggregateShareMetadataReadsPerSec    *prometheus.Desc
	AggregateUsecPerFileMetadataCreateOp *prometheus.Desc
	AggregateUsecPerFileMetadataModifyOp *prometheus.Desc
	AggregateUsecPerFileMetadataReadOp   *prometheus.Desc
	AggregateUsecPerShareMetadataReadOp  *prometheus.Desc
	CreatesPerSec                        *prometheus.Desc
	FsinfosPerSec                        *prometheus.Desc
	FsstatsPerSec                        *prometheus.Desc
	GetattrsPerSec                       *prometheus.Desc
	LinksPerSec                          *prometheus.Desc
	LookupsPerSec                        *prometheus.Desc
	MkdirsPerSec                         *prometheus.Desc
	OthersPerSec                         *prometheus.Desc
	PathconfsPerSec                      *prometheus.Desc
	ReadsPerSec                          *prometheus.Desc
	ReaddirsPerSec                       *prometheus.Desc
	ReaddirplusesPerSec                  *prometheus.Desc
	ReadlinksPerSec                      *prometheus.Desc
	RemovesPerSec                        *prometheus.Desc
	RenamesPerSec                        *prometheus.Desc
	RmdirsPerSec                         *prometheus.Desc
	SetattrsPerSec                       *prometheus.Desc
	SymlinksPerSec                       *prometheus.Desc
	WritesPerSec                         *prometheus.Desc
	UsecPerAccessOp                      *prometheus.Desc
	UsecPerCreateOp                      *prometheus.Desc
	UsecPerFsinfoOp                      *prometheus.Desc
	UsecPerFsstatOp                      *prometheus.Desc
	UsecPerGetattrOp                     *prometheus.Desc
	UsecPerLinkOp                        *prometheus.Desc
	UsecPerLookupOp                      *prometheus.Desc
	UsecPerMkdirOp                       *prometheus.Desc
	UsecPerOtherOp                       *prometheus.Desc
	UsecPerPathconfOp                    *prometheus.Desc
	UsecPerReadOp                        *prometheus.Desc
	UsecPerReaddirOp                     *prometheus.Desc
	UsecPerReaddirPlusOp                 *prometheus.Desc
	UsecPerReadlinkOp                    *prometheus.Desc
	UsecPerRemoveOp                      *prometheus.Desc
	UsecPerRenameOp                      *prometheus.Desc
	UsecPerRmdirOp                       *prometheus.Desc
	UsecPerSetattrOp                     *prometheus.Desc
	UsecPerSymlinkOp                     *prometheus.Desc
	UsecPerWriteOp                       *prometheus.Desc
}

func NewArrayNfsPerformanceCollector(fbClient *fb.FlashbladeClient) *ArrayNfsPerformanceCollector {
	const subsystem = "perf_nfs"

	return &ArrayNfsPerformanceCollector{
		fbClient: fbClient,
		AccessesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "accesses_per_sec"),
			"ACCESS requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateFileMetadataCreatesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_file_metadata_creates_per_sec"),
			"Sum of file-level or directory-level create-like metadata requests per second. Includes CREATE, LINK, MKDIR, and SYMLINK",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateFileMetadataModifiesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_file_metadata_modifies_per_sec"),
			"Sum of file-level or directory-level modify-like and delete-like metadata requests per second. Includes REMOVE, RENAME, RMDIR, and SETATTR",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateFileMetadataReadsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_file_metadata_reads_per_sec"),
			"Sum of file-level or directory-level read-like metadata requests per second. Includes GETATTR, LOOKUP, PATHCONF, READDIR, READDIRPLUS, and READLINK",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateShareMetadataReadsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_share_metadata_reads_per_sec"),
			"Sum of share-level read-like metadata requests per second. Includes ACCESS, FSINFO, and FSSTAT",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateUsecPerFileMetadataCreateOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_usec_per_file_metadata_create_op"),
			"Average time, measured in microseconds, it takes the array to process a file-level or directory-level create-like metadata request. Includes CREATE, LINK, MKDIR, and SYMLINK",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateUsecPerFileMetadataModifyOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_usec_per_file_metadata_modify_op"),
			"Average time, measured in microseconds, it takes the array to process a file-level or directory-level modify-like or delete-like metadata request. Includes REMOVE, RENAME, RMDIR, and SETATTR",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateUsecPerFileMetadataReadOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_usec_per_file_metadata_read_op"),
			"Average time, measured in microseconds, it takes the array to process a file-level or directory-level read-like metadata request. Includes GETATTR, LOOKUP, PATHCONF, READDIR, READDIRPLUS, and READLINK",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		AggregateUsecPerShareMetadataReadOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "aggregate_usec_per_share_metadata_read_op"),
			"Average time, measured in microseconds, it takes the array to process a share-level read-like metadata request. Includes ACCESS, FSINFO, and FSSTAT",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		CreatesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "creates_per_sec"),
			"CREATE requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		FsinfosPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "fsinfos_per_sec"),
			"FSINFO requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		FsstatsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "fsstats_per_sec"),
			"FSSTAT requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		GetattrsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "getattrs_per_sec"),
			"GETATTR requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		LinksPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "links_per_sec"),
			"LINK requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		LookupsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "lookups_per_sec"),
			"LOOKUP requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		MkdirsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "mkdirs_per_sec"),
			"MKDIR requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		OthersPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "others_per_sec"),
			"All other requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		PathconfsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "pathconfs_per_sec"),
			"PATHCONF requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reads_per_sec"),
			"READ requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReaddirsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "readdirs_per_sec"),
			"READDIR requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReaddirplusesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "readdirpluses_per_sec"),
			"READDIRPLUS requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		ReadlinksPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "readlinks_per_sec"),
			"READLINK requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		RemovesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "removes_per_sec"),
			"REMOVE requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		RenamesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "renames_per_sec"),
			"RENAME requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		RmdirsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "rmdirs_per_sec"),
			"RMDIR requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		SetattrsPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "setattrs_per_sec"),
			"SETATTR requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		SymlinksPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "symlinks_per_sec"),
			"SYMLINK requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		WritesPerSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "writes_per_sec"),
			"WRITE requests processed per second",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerAccessOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_access_op"),
			"Average time, measured in microseconds, it takes the array to process an ACCESS request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerCreateOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_create_op"),
			"Average time, measured in microseconds, it takes the array to process a CREATE request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerFsinfoOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_fsinfo_op"),
			"Average time, measured in microseconds, it takes the array to process an FSINFO request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerFsstatOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_fsstat_op"),
			"Average time, measured in microseconds, it takes the array to process an FSSTAT request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerGetattrOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_getattr_op"),
			"Average time, measured in microseconds, it takes the array to process a GETATTR request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerLinkOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_link_op"),
			"Average time, measured in microseconds, it takes the array to process a LINK request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerLookupOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_lookup_op"),
			"Average time, measured in microseconds, it takes the array to process a LOOKUP request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerMkdirOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_mkdir_op"),
			"Average time, measured in microseconds, it takes the array to process a MKDIR request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerOtherOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_other_op"),
			"Average time, measured in microseconds, it takes the array to process all other requests",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerPathconfOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_pathconf_op"),
			"Average time, measured in microseconds, it takes the array to process a PATHCONF request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_read_op"),
			"Average time, measured in microseconds, it takes the array to process a READ request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReaddirOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_readdir_op"),
			"Average time, measured in microseconds, it takes the array to process a READDIR request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReaddirPlusOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_readdir_plus_op"),
			"Average time, measured in microseconds, it takes the array to process a READDIRPLUS request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerReadlinkOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_readlink_op"),
			"Average time, measured in microseconds, it takes the array to process a READLINK request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerRemoveOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_remove_op"),
			"Average time, measured in microseconds, it takes the array to process a REMOVE request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerRenameOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_rename_op"),
			"Average time, measured in microseconds, it takes the array to process a RENAME request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerRmdirOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_rmdir_op"),
			"Average time, measured in microseconds, it takes the array to process an RMDIR request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerSetattrOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_setattr_op"),
			"Average time, measured in microseconds, it takes the array to process a SETATTR request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerSymlinkOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_symlink_op"),
			"Average time, measured in microseconds, it takes the array to process a SYMLINK request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
		UsecPerWriteOp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "usec_per_write_op"),
			"Average time, measured in microseconds, it takes the array to process a WRITE request",
			nil,
			prometheus.Labels{"host": fbClient.Host},
		),
	}
}

func (c ArrayNfsPerformanceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Error("Failed collecting array NFS performance metrics", err)
	}
}

func (c ArrayNfsPerformanceCollector) collect(ch chan<- prometheus.Metric) error {
	stats, err := c.fbClient.ArrayNfsPerformance()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		c.AccessesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].AccessesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateFileMetadataCreatesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateFileMetadataCreatesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateFileMetadataModifiesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateFileMetadataModifiesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateFileMetadataReadsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateFileMetadataReadsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateShareMetadataReadsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateShareMetadataReadsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateUsecPerFileMetadataCreateOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateUsecPerFileMetadataCreateOp))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateUsecPerFileMetadataModifyOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateUsecPerFileMetadataModifyOp))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateUsecPerFileMetadataReadOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateUsecPerFileMetadataReadOp))

	ch <- prometheus.MustNewConstMetric(
		c.AggregateUsecPerShareMetadataReadOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].AggregateUsecPerShareMetadataReadOp))

	ch <- prometheus.MustNewConstMetric(
		c.CreatesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].CreatesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.FsinfosPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].FsinfosPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.FsstatsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].FsstatsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.GetattrsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].GetattrsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.LinksPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].LinksPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.LookupsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].LookupsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.MkdirsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].MkdirsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.OthersPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].OthersPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.PathconfsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].PathconfsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReadsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].CreatesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReaddirsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReaddirsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReaddirplusesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReaddirplusesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.ReadlinksPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].ReadlinksPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.RemovesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].RemovesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.RenamesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].RenamesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.RmdirsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].RmdirsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.SetattrsPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].SetattrsPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.SymlinksPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].SymlinksPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.WritesPerSec,
		prometheus.GaugeValue,
		float64(stats.Items[0].WritesPerSec))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerAccessOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerAccessOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerCreateOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerCreateOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerFsinfoOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerFsinfoOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerFsstatOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerFsstatOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerGetattrOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerGetattrOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerLinkOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerLinkOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerLookupOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerLookupOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerMkdirOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerMkdirOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerOtherOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerOtherOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerPathconfOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerPathconfOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReaddirOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReaddirOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReaddirPlusOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReaddirPlusOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerReadlinkOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerReadlinkOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerRemoveOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerRemoveOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerRenameOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerRenameOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerRmdirOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerRmdirOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerSetattrOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerSetattrOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerSymlinkOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerSymlinkOp))

	ch <- prometheus.MustNewConstMetric(
		c.UsecPerWriteOp,
		prometheus.GaugeValue,
		float64(stats.Items[0].UsecPerWriteOp))

	return nil
}
