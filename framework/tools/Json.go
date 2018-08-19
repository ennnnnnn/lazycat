package tools

import (
	"encoding/json"
	"fmt"
)

func Map2Json(data map[string]interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return []byte("")
	}
	return b
}
func Object2Json(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return []byte("")
	}
	return b
}