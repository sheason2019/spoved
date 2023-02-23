package utils

import "encoding/json"

func StringToMap(str string) map[string]string {
	m := map[string]string{}
	json.Unmarshal([]byte(str), &m)

	return m
}
