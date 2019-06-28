// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type BladesResponse struct {
	Items []BladesItem `json:"items"`
}

type BladesItem struct {
	Name        string  `json:"name"`
	Details     string  `json:"details"`
	RawCapacity int     `json:"raw_capacity"`
	Target      string  `json:"target"`
	Progress    float64 `json:"progress"`
	Status      string  `json:"status"`
}

func (fbClient FlashbladeClient) Blades() (BladesResponse, error) {
	endpoint := "/1.2/blades"
	var bladesResponse BladesResponse
	err := fbClient.GetJSON(endpoint, nil, &bladesResponse)
	return bladesResponse, err
}
