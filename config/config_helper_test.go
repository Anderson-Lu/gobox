package config

import (
	"reflect"
	"testing"
)

type MyConfig struct {
	Name   string  `yaml:"name"`
	Value  string  `yaml:"value"`
	Int    int     `yaml:"int"`
	Double float64 `yaml:"double"`
	T      T       `yaml:"t"`
}

type T struct {
	Name string `yaml:"tname"`
}

func TestInitConfigFile(t *testing.T) {
	var myconf = MyConfig{}

	var expect = MyConfig{
		Name:   "helloword",
		Value:  "i am value",
		Int:    1,
		Double: 1.11,
		T: T{
			Name: "t",
		},
	}

	type args struct {
		path   string
		format int
		value  MyConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试加载json配置文件",
			args: args{
				path:   "./config_test.json",
				format: CONFIG_JSON_FORMAT,
				value:  myconf,
			},
			wantErr: false,
		},
		{
			name: "测试加载yaml配置文件",
			args: args{
				path:   "./config_test.yaml",
				format: CONFIG_YAML_FORMAT,
				value:  myconf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfigFile(tt.args.path, tt.args.format, &tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InitConfigFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.value, expect) {
				t.Errorf("InitConfigFile() error =%v, want = %v", tt.args.value, expect)
			}
		})
	}
}
