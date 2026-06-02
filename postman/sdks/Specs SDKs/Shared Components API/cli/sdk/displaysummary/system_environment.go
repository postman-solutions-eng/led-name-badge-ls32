package displaysummary

import "encoding/json"

// The API catalog system environment used to list services
type SystemEnvironment struct {
	ID   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SystemEnvironment) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemEnvironment to string"
	}
	return string(jsonData)
}
