package displaysummary

import "encoding/json"

type CreateDisplaySummaryOkResponse struct {
	// Status of the operation
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Present when a Postman API key was supplied. Indicates the summary was built from the Postman API catalog.
	Source *Source `json:"source,omitempty" xml:"source,omitempty"`
	// The text rendered on the LED badge
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The API catalog system environment used to list services
	SystemEnvironment *SystemEnvironment `json:"systemEnvironment,omitempty" xml:"systemEnvironment,omitempty"`
	// Number of services returned for the summary
	ServiceCount *int64 `json:"serviceCount,omitempty" xml:"serviceCount,omitempty"`
	// True when additional catalog services exist beyond the fetched page
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// Catalog services included in the summary
	Services []Services `json:"services,omitempty" xml:"services,omitempty"`
}

func (c CreateDisplaySummaryOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateDisplaySummaryOkResponse to string"
	}
	return string(jsonData)
}
