package fb

type ArrayPerformanceResponse struct {
	Items [1]ArrayPerformanceItem `json:"items"`
}

type ArrayPerformanceItem struct {
	Name           string  `json:"name"`
	BytesPerOp     float64 `json:"bytes_per_op"`
	BytesPerRead   float64 `json:"bytes_per_read"`
	BytesPerWrite  float64 `json:"bytes_per_write"`
	OthersPerSec   float64 `json:"others_per_sec"`
	OutputPerSec   float64 `json:"output_per_sec"`
	ReadsPerSec    float64 `json:"reads_per_sec"`
	Time           float64 `json:"time"`
	UsecPerOtherOp float64 `json:"usec_per_other_op"`
	UsecPerReadOp  float64 `json:"usec_per_read_op"`
	UsecPerWriteOp float64 `json:"usec_per_write_op"`
	InputPerSec    float64 `json:"input_per_sec"`
	WritesPerSec   float64 `json:"writes_per_sec"`
}

func (fbClient FlashbladeClient) ArrayPerformance() (ArrayPerformanceResponse, error) {
	endpoint := "/1.2/arrays/performance"
	var arrayPerformanceResponse ArrayPerformanceResponse
	err := fbClient.GetJSON(endpoint, nil, &arrayPerformanceResponse)
	return arrayPerformanceResponse, err
}
