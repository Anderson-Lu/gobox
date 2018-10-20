package config_helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const CONFIG_JSON_FORMAT int = 0

//加载配置信息
func InitConfigFile(path string, format int, value interface{}) error {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	switch format {
	case CONFIG_JSON_FORMAT:
		json.Unmarshal(file, &value)
	default:
		return fmt.Errorf("invalid format type: %d", format)
	}

	return nil
}
