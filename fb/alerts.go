// Copyright (C) 2019 by the authors in the project README.md
// See the full license in the project LICENSE file.

package fb

type AlertsResponse struct {
	Items []AlertsItem `json:"items"`
}

type AlertsItem struct {
	Name             string            `json:"name"`
	Created          int               `json:"created"`
	Index            int               `json:"index"`
	Code             int               `json:"code"`
	Severity         string            `json:"severity"`
	Component        string            `json:"component"`
	State            string            `json:"state"`
	Flagged          bool              `json:"flagged"`
	Updated          int               `json:"updated"`
	Notified         int               `json:"notified"`
	Subject          string            `json:"subject"`
	Description      string            `json:"description"`
	KnowledgeBaseURL string            `json:"knowledge_base_url"`
	Acion            string            `json:"action"`
	Variables        map[string]string `json:"variables"`
}

func (fbClient FlashbladeClient) OpenAlerts() (AlertsResponse, error) {
	endpoint := "/1.2/alerts"
	var alertsResponse AlertsResponse
	params := make(map[string]string)
	params["filter"] = "state='open'"
	err := fbClient.GetJSON(endpoint, params, &alertsResponse)
	return alertsResponse, err
}
