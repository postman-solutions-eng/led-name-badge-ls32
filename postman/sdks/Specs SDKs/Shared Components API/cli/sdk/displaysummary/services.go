package displaysummary

import "encoding/json"

type Services struct {
	ID     *string `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *Status `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Services) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Services to string"
	}
	return string(jsonData)
}
