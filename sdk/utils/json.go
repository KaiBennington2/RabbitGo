package utils

import "encoding/json"

func ToJson(e any) ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return data, nil
}
