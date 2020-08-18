// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type ArrayS3PerformanceResponse struct {
	Items [1]ArrayS3PerformanceItem `json:"items"`
}

type ArrayS3PerformanceItem struct {
	Name                 string  `json:"name"`
	ReadBucketsPerSec    float64 `json:"read_buckets_per_sec"`
	WriteBucketsPerSec   float64 `json:"write_buckets_per_sec"`
	ReadObjectsPerSec    float64 `json:"read_objects_per_sec"`
	WriteObjectsPerSec   float64 `json:"write_objects_per_sec"`
	OthersPerSec         float64 `json:"others_per_sec"`
	UsecPerReadBucketOp  float64 `json:"usec_per_read_bucket_op"`
	UsecPerWriteBucketOp float64 `json:"usec_per_write_bucket_op"`
	UsecPerReadObjectOp  float64 `json:"usec_per_read_object_op"`
	UsecPerWriteObjectOp float64 `json:"usec_per_write_object_op"`
	UsecPerOtherOp       float64 `json:"usec_per_other_op"`
	Time                 float64 `json:"time"`
}

func (fbClient FlashbladeClient) ArrayS3Performance() (ArrayS3PerformanceResponse, error) {
	endpoint := "arrays/s3-specific-performance"
	var arrayS3PerformanceResponse ArrayS3PerformanceResponse
	err := fbClient.GetJSON(endpoint, nil, &arrayS3PerformanceResponse)
	return arrayS3PerformanceResponse, err
}
