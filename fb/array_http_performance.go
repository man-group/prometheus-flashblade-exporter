// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type ArrayHttpPerformanceResponse struct {
	Items [1]ArrayHttpPerformanceItem `json:"items"`
}

type ArrayHttpPerformanceItem struct {
	Name               string  `json:"name"`
	ReadDirsPerSec     float64 `json:"read_dirs_per_sec"`
	WriteDirsPerSec    float64 `json:"write_dirs_per_sec"`
	ReadFilesPerSec    float64 `json:"read_files_per_sec"`
	WriteFilesPerSec   float64 `json:"write_files_per_sec"`
	OthersPerSec       float64 `json:"others_per_sec"`
	UsecPerReadDirOp   float64 `json:"usec_per_read_dir_op"`
	UsecPerWriteDirOp  float64 `json:"usec_per_write_dir_op"`
	UsecPerReadFileOp  float64 `json:"usec_per_read_file_op"`
	UsecPerWriteFileOp float64 `json:"usec_per_write_file_op"`
	UsecPerOtherOp     float64 `json:"usec_per_other_op"`
	Time               float64 `json:"time"`
}

func (fbClient FlashbladeClient) ArrayHttpPerformance() (ArrayHttpPerformanceResponse, error) {
	endpoint := "arrays/http-specific-performance"
	var arrayHttpPerformanceResponse ArrayHttpPerformanceResponse
	err := fbClient.GetJSON(endpoint, nil, &arrayHttpPerformanceResponse)
	return arrayHttpPerformanceResponse, err
}
