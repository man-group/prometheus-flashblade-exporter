// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type FSPerformanceResponse struct {
	Items []FSPerformanceItem `json:"items"`
}

type FSPerformanceItem struct {
	Name             string  `json:"name"`
	BytesPerOp       float64 `json:"bytes_per_op"`
	BytesPerRead     float64 `json:"bytes_per_read"`
	BytesPerWrite    float64 `json:"bytes_per_write"`
	OthersPerSec     float64 `json:"others_per_sec"`
	ReadBytesPerSec  float64 `json:"read_bytes_per_sec"`
	ReadsPerSec      float64 `json:"reads_per_sec"`
	Time             float64 `json:"time"`
	UsecPerOtherOp   float64 `json:"usec_per_other_op"`
	UsecPerReadOp    float64 `json:"usec_per_read_op"`
	UsecPerWriteOp   float64 `json:"usec_per_write_op"`
	WritesPerSec     float64 `json:"writes_per_sec"`
	WriteBytesPerSec float64 `json:"write_bytes_per_sec"`
}

type FS struct {
	FileSystem struct {
		Name         string `json:"name"`
		Id           string `json:"id"`
		ResourceType string `json:"resource_type"`
	} `json:"file_system"`
}

type FSUser struct {
	User struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	} `json:"user"`
}

type FSUserPerformanceResp struct {
	Items []struct {
		FSPerformanceItem
		FS
		FSUser
	} `json:"items"`
}

func (fbClient FlashbladeClient) FSPerformance() (FSPerformanceResponse, error) {
	endpoint := "file-systems/performance"

	params := make(map[string]string)
	params["protocol"] = "nfs" // Only NFS supported as of API 1.8

	var fsPerformance FSPerformanceResponse
	err := fbClient.GetJSON(endpoint, params, &fsPerformance)
	return fsPerformance, err
}

// FSPerformanceByUser returns per user, per file system NFS counters
func (fbClient FlashbladeClient) FSPerformanceByUser() ([]FSUserPerformanceResp, error) {
	fsU := []FSUserPerformanceResp{}
	// we first need to get the list of file systems
	filesystems, err := fbClient.Filesystems()
	if err != nil {
		return fsU, err
	}

	endpoint := "file-systems/users/performance"

	for _, f := range filesystems.Items {
		params := make(map[string]string)
		params["protocol"] = "nfs" // Only NFS supported as of API 1.8
		params["file_system_names"] = f.Name
		fs := &FSUserPerformanceResp{}

		err := fbClient.GetJSON(endpoint, params, fs)
		if err != nil {
			return fsU, err
		}
		fsU = append(fsU, *fs)
	}

	return fsU, nil
}
