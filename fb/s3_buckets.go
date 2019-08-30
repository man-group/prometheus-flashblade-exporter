// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type S3BucketsResponse struct {
	Total BucketItem   `json:"total"`
	Items []BucketItem `json:"items"`
}

type BucketItem struct {
	Name          string `json:"name"`
	Created       int    `json:"created"`
	Space         Space  `json:"space"`
	Destroyed     bool   `json:"destroyed"`
	TimeRemaining int    `json:"time_remaining"`
	ObjectCount   int    `json:"object_count"`
	Versioning    string `json:"versioning"`
}

func (fbClient FlashbladeClient) S3Buckets() (S3BucketsResponse, error) {
	endpoint := "buckets"
	var s3BucketsResponse S3BucketsResponse
	err := fbClient.GetJSON(endpoint, nil, &s3BucketsResponse)
	return s3BucketsResponse, err
}
