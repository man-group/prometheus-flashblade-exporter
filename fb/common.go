// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type Space struct {
	Virtual       int     `json:"virtual"`
	DataReduction float64 `json:"data_reduction"`
	Unique        int     `json:"unique"`
	Snapshots     int     `json:"snapshots"`
	TotalPhysical int     `json:"total_physical"`
}
