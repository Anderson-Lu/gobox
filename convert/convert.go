package convert

import "encoding/json"

func ConvertBs2Interface(bs []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	e := json.Unmarshal(bs, &result)
	return result, e
}
