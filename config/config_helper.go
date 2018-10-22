package config_helper

import (
	"reflect"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_JSON_FORMAT int = 0
	CONFIG_YAML_FORMAT int = 1
)

//加载配置信息
func InitConfigFile(path string, format int, value interface{}) error {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	switch format {
	case CONFIG_JSON_FORMAT:
		return json.Unmarshal(file, &value)
	case CONFIG_YAML_FORMAT:
		e := yaml.Unmarshal(file, value)
		fmt.Println(e, value,reflect.TypeOf(value))
		return e
	default:
		return fmt.Errorf("invalid format type: %d", format)
	}
}
