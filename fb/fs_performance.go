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

func (fbClient FlashbladeClient) FSPerformance() (FSPerformanceResponse, error) {
	endpoint := "/1.8/file-systems/performance"
	params := make(map[string]string)
	params["protocol"] = "nfs" // Only NFS supported as of API 1.8

	var fsPerformance FSPerformanceResponse
	err := fbClient.GetJSON(endpoint, params, &fsPerformance)
	return fsPerformance, err
}
