// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type ArraySpaceResponse struct {
	Items [1]ArraySpaceItem `json:"items"`
}

type ArraySpaceItem struct {
	Name     string  `json:"name"`
	Capacity float64 `json:"capacity"`
	Parity   float64 `json:"parity"`
	Space    Space   `json:"space"`
	Time     float64 `json:"time"`
}

func (fbClient FlashbladeClient) ArraySpace() (ArraySpaceResponse, error) {
	endpoint := "/1.2/arrays/space"
	var arraySpaceResponse ArraySpaceResponse
	err := fbClient.GetJSON(endpoint, nil, &arraySpaceResponse)
	return arraySpaceResponse, err
}
