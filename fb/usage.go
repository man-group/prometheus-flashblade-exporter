// Copyright (c) 2019 Hudson River Trading LLC
// All rights reserved.

package fb

import "github.com/prometheus/common/log"

type UsageResponse struct {
	Groups []UsageGroup
	Users  []UsageUser
}

type PaginationData struct {
	ContinuationToken string `json:"continuation_token"`
	Total             int    `json:"total_item_count"`
}

type UsageGroup struct {
	Items          []UsageItemGroup `json:"items"`
	PaginationInfo PaginationData   `json:"pagination_info"`
}

type UsageUser struct {
	Items          []UsageItemUser `json:"items"`
	PaginationInfo PaginationData  `json:"pagination_info"`
}

type UsageItemGroup struct {
	FileSystem             map[string]string `json:"file_system"`
	FileSystemDefaultQuota int               `json:"file_system_default_quota"`
	Group                  NameID            `json:"group"`
	Name                   string            `json:"name"`
	Quota                  int               `json:"quota"`
	Usage                  int               `json:"usage"`
}

type UsageItemUser struct {
	FileSystem             map[string]string `json:"file_system"`
	FileSystemDefaultQuota int               `json:"file_system_default_quota"`
	User                   NameID            `json:"user"`
	Name                   string            `json:"name"`
	Quota                  int               `json:"quota"`
	Usage                  int               `json:"usage"`
}

type NameID struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (fbClient FlashbladeClient) Usage() (UsageResponse, error) {
	endpoint := "1.8/file-systems"
	var filesystemsResponse FilesystemsResponse
	err := fbClient.GetJSON(endpoint, nil, &filesystemsResponse)

	if err != nil {
		log.Errorf("Error while getting JSON on endpoint %s", endpoint, err)
		return UsageResponse{}, err
	}

	var (
		usageResponseGroup []UsageGroup
		usageResponseUser  []UsageUser
		usageGroup UsageGroup
		usageUser UsageUser
	)
	params := make(map[string]string)

	for _, item := range filesystemsResponse.Items {
		params["file_system_names"] = item.Name

		endpoint = "1.8/usage/groups"
		err = fbClient.GetJSON(endpoint, params, &usageGroup)
		if err != nil {
			log.Errorf("Error while getting JSON on endpoint %s", endpoint, err)
			return UsageResponse{}, err
		}
		usageResponseGroup = append(usageResponseGroup, usageGroup)

		endpoint = "1.8/usage/users"
		err = fbClient.GetJSON(endpoint, params, &usageUser)
		if err != nil {
			log.Errorf("Error while getting JSON on endpoint %s", endpoint, err)
			return UsageResponse{}, err
		}
		
		usageResponseUser = append(usageResponseUser, usageUser)

	}

	return UsageResponse{usageResponseGroup, usageResponseUser}, nil
}
