package method

import "encoding/json"

func GoBytes(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return bytes
}
