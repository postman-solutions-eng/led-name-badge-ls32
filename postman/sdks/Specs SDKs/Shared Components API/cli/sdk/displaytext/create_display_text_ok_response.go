package displaytext

import "encoding/json"

type CreateDisplayTextOkResponse struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Text   *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (c CreateDisplayTextOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateDisplayTextOkResponse to string"
	}
	return string(jsonData)
}
