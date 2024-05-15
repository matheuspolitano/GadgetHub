package chat

import "encoding/json"

func JSONToMap(jsonStr string) (map[string]string, error) {
	result := make(map[string]string)

	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func MapToJSON(data map[string]string) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
