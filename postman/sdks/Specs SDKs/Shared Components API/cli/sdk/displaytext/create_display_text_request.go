package displaytext

import "encoding/json"

type CreateDisplayTextRequest struct {
	// The text to display, optionally including icon codes in icon_name format
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (c CreateDisplayTextRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateDisplayTextRequest to string"
	}
	return string(jsonData)
}
