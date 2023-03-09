package config

import (
	"log"
	"reflect"
	"testing"
)

// 测试用的配置文件
const ConfigFilePath = "./config_test.yaml"

func init() {
	cfg, err := Load(ConfigFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	SetGlobalConfig(cfg)
}

func TestGetString(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want string
		rsp  string
	}{
		{
			"demo_string",
			"this is a string",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetString(tt.name, "default")
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want int
		rsp  int
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetInt(tt.name, -1)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetInt32(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want int32
		rsp  int32
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetInt32(tt.name, -1)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetInt64(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want int64
		rsp  int64
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetInt64(tt.name, -1)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetUint(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want uint
		rsp  uint
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetUint(tt.name, 0)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetUint32(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want uint32
		rsp  uint32
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetUint32(tt.name, 0)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetUint64(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want uint64
		rsp  uint64
	}{
		{
			"demo_int",
			1234567,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetUint64(tt.name, 0)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want bool
		rsp  bool
	}{
		{
			"demo_bool",
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetBool(tt.name, false)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetFloat32(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want float32
		rsp  float32
	}{
		{
			"demo_float",
			0.3214,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetFloat32(tt.name, 0)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want float64
		rsp  float64
	}{
		{
			"demo_float",
			0.3214,
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetFloat64(tt.name, 0)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestDecodeNormal(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want any
		rsp  any
	}{
		{
			"demo_string",
			"this is a string",
			"",
		},
		{
			"demo_int",
			1234567,
			0,
		},
		{
			"demo_bool",
			true,
			false,
		},
		{
			"demo_float",
			0.3214,
			false,
		},
		{
			"demo_array",
			[]any{
				"array01",
				"array02",
				"array03",
			},
			[]string{},
		},
		{
			"demo_map",
			map[string]any{
				"map01": "some message",
				"map02": "other message",
			},
			map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cfg.Decode(tt.name, &tt.rsp); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}

}

func TestMultiLevelGet(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		want any
		rsp  any
	}{
		{
			"demo_multi_level.level1.level2.level3.level4.level5.level6.level7.level8.level9.level10",
			"This is level 10",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cfg.Decode(tt.name, &tt.rsp); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestDecodeStruct(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	type Person struct {
		Name    string `yaml:"name"`
		Address struct {
			Street string `yaml:"street"`
			City   string `yaml:"city"`
			State  string `yaml:"state"`
			Zip    []int  `yaml:"zip"`
		}
	}
	tests := []struct {
		name string
		want []Person
		rsp  []Person
	}{
		{
			name: "demo_struct",
			want: []Person{
				{
					"John Doe",
					struct {
						Street string `yaml:"street"`
						City   string `yaml:"city"`
						State  string `yaml:"state"`
						Zip    []int  `yaml:"zip"`
					}{
						"123 E 3rd St",
						"Denver",
						"CO",
						[]int{81526, 12345},
					},
				},
				{
					"Torben Ad",
					struct {
						Street string `yaml:"street"`
						City   string `yaml:"city"`
						State  string `yaml:"state"`
						Zip    []int  `yaml:"zip"`
					}{
						"皇后大街189号",
						"香港",
						"CN",
						[]int{88675},
					},
				},
			},
			rsp: []Person{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cfg.Decode(tt.name, &tt.rsp); err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestLoadAndDecode(t *testing.T) {

	type Person struct {
		Name    string `yaml:"name"`
		Address struct {
			Street string `yaml:"street"`
			City   string `yaml:"city"`
			State  string `yaml:"state"`
			Zip    []int  `yaml:"zip"`
		}
	}
	type AllConfig struct {
		DemoString string            `yaml:"demo_string"`
		DemoInt    int               `yaml:"demo_int"`
		DemoBool   bool              `yaml:"demo_bool"`
		DemoFloat  float64           `yaml:"demo_float"`
		DemoArray  []string          `yaml:"demo_array"`
		DemoMap    map[string]string `yaml:"demo_map"`
		DemoStruct []Person          `yaml:"demo_struct"`
	}

	tests := []struct {
		name string
		want AllConfig
		rsp  AllConfig
	}{
		{
			"test decode all config",
			AllConfig{
				DemoString: "this is a string",
				DemoInt:    1234567,
				DemoBool:   true,
				DemoFloat:  0.3214,
				DemoArray:  []string{"array01", "array02", "array03"},
				DemoMap:    map[string]string{"map01": "some message", "map02": "other message"},
				DemoStruct: []Person{
					{
						Name: "John Doe",
						Address: struct {
							Street string `yaml:"street"`
							City   string `yaml:"city"`
							State  string `yaml:"state"`
							Zip    []int  `yaml:"zip"`
						}{Street: "123 E 3rd St", City: "Denver", State: "CO", Zip: []int{81526, 12345}},
					},
					{
						Name: "Torben Ad",
						Address: struct {
							Street string `yaml:"street"`
							City   string `yaml:"city"`
							State  string `yaml:"state"`
							Zip    []int  `yaml:"zip"`
						}{Street: "皇后大街189号", City: "香港", State: "CN", Zip: []int{88675}},
					},
				},
			},
			AllConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadAndDecode(ConfigFilePath, &tt.rsp); err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestNoKey01(t *testing.T) {
	cfg, _ := GetGlobalConfig().(*Config)
	tests := []struct {
		name string
		key  string
		want string
		rsp  string
	}{
		{
			"have no key",
			"",
			"default",
			"",
		},
		{
			"key is not exist",
			"demo_multi_level.level1.level2.level3.level4.level5.level6.level7.level8.level9.level10.level11",
			"default",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rsp = cfg.GetString(tt.key, tt.want)
			if !reflect.DeepEqual(tt.rsp, tt.want) {
				t.Errorf("%s got = %v, want %v", tt.name, tt.rsp, tt.want)
			}
		})
	}
}

func TestFileNotExist(t *testing.T) {

	tests := []struct {
		name      string
		filePath  string
		wantError bool
	}{
		{
			"config file exist",
			ConfigFilePath,
			false,
		},
		{
			"config file is not exist",
			"./file_is_not_exist.yaml",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Load(tt.filePath); err != nil && tt.wantError != true {
				t.Errorf("%s got err %v", tt.name, err)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			var v any
			if err := LoadAndDecode(tt.filePath, &v); err != nil && tt.wantError != true {
				t.Errorf("%s got err %v", tt.name, err)
			}
		})
	}
}
