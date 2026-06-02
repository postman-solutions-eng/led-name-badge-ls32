package predefinedicons

import "encoding/json"

type GetPredefinedIconsOkResponse struct {
	// Array of icon codes in :icon_name format
	Icons []string `json:"icons,omitempty" xml:"icons,omitempty"`
}

func (g GetPredefinedIconsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPredefinedIconsOkResponse to string"
	}
	return string(jsonData)
}
