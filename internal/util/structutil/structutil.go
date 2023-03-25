package structutil

import "encoding/json"

func ToJsonString(data interface{}) string {
	jsonString, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return err.Error()
	}

	return string(jsonString)
}
