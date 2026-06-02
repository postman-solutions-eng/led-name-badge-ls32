package displaysummary

import "encoding/json"

type CreateDisplaySummaryRequest struct {
	// The type of summary message to display (e.g., welcome, status, alert, info)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Optional custom text to append to the summary
	CustomText *string `json:"customText,omitempty" xml:"customText,omitempty"`
	// Optional alternative to the `X-API-Key` header for supplying a Postman API key (`PMAK-...`) used to fetch the API catalog
	APIKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// Optional API catalog system environment ID. When omitted, the first production environment is used, otherwise the first environment with associations, otherwise the first available environment.
	SystemEnvironmentID *string `json:"systemEnvironmentId,omitempty" xml:"systemEnvironmentId,omitempty"`
}

func (c CreateDisplaySummaryRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateDisplaySummaryRequest to string"
	}
	return string(jsonData)
}
