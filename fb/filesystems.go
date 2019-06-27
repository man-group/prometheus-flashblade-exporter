// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type FilesystemsResponse struct {
	Total FilesystemsItem   `json:"total"`
	Items []FilesystemsItem `json:"items"`
}

type FilesystemsItem struct {
	Name                       string      `json:"name"`
	Created                    int         `json:"created"`
	FastRemoveDirectoryEnabled bool        `json:"fast_remove_directory_enabled"`
	Provisioned                int         `json:"target"`
	SnapshotDirectoryEnabled   bool        `json:"progress"`
	Space                      Space       `json:"space"`
	NFS                        interface{} `json:"nfs"`
	HTTP                       interface{} `json:"http"`
	SMB                        interface{} `json:"smb"`
	Destroyed                  bool        `json:"destroyed"`
	TimeRemaining              int         `json:"time_remaining"`
}

type Space struct {
	Virtual       int     `json:"virtual"`
	DataReduction float64 `json:"data_reduction"`
	Unique        int     `json:"unique"`
	Snapshots     int     `json:"snapshots"`
	TotalPhysical int     `json:"total_physical"`
}

func (fbClient FlashbladeClient) Filesystems() (FilesystemsResponse, error) {
	endpoint := "/1.2/file-systems"
	var filesystemsResponse FilesystemsResponse
	err := fbClient.GetJSON(endpoint, nil, &filesystemsResponse)
	return filesystemsResponse, err
}
