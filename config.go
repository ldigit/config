package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync/atomic"
)

// SEPARATOR 配置名称中的层级分隔符
const SEPARATOR = "."

var globalConfig atomic.Value

type Config map[string]yaml.Node

// SetGlobalConfig 存储到全局配置
func SetGlobalConfig(cfg any) {
	globalConfig.Store(cfg)
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() any {
	return globalConfig.Load()
}

// Load 加载配置文件内容
func Load(path string) (*Config, error) {
	// 判断文件是否存在
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return nil, err
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// LoadAndDecode 加载配置并解析到对象
func LoadAndDecode(path string, v interface{}) error {
	// 判断文件是否存在
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return err
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	if err := yaml.Unmarshal(buf, v); err != nil {
		return err
	}

	return nil
}

// Decode 获取指定KEY的结果值
func (c Config) Decode(key string, v any) error {
	keys := strings.Split(key, SEPARATOR)

	if 0 == len(keys) {
		return errors.New("key is empty")
	}

	n, err := find(keys, c)
	if err != nil {
		return err
	}

	return n.Decode(v)
}

// GetString 获取指定KEY的字符串值
func (c Config) GetString(key string, defaultValue string) string {
	var v string
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetBool 获取指定KEY的Bool值
func (c Config) GetBool(key string, defaultValue bool) bool {
	var v bool
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetInt 获取指定KEY的Int值
func (c Config) GetInt(key string, defaultValue int) int {
	var v int
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetInt32 获取指定KEY的Int32值
func (c Config) GetInt32(key string, defaultValue int32) int32 {
	var v int32
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetInt64 获取指定KEY的Int64值
func (c Config) GetInt64(key string, defaultValue int64) int64 {
	var v int64
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetUint 获取指定KEY的Uint值
func (c Config) GetUint(key string, defaultValue uint) uint {
	var v uint
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetUint32 获取指定KEY的Uint32值
func (c Config) GetUint32(key string, defaultValue uint32) uint32 {
	var v uint32
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetUint64 获取指定KEY的Uint64值
func (c Config) GetUint64(key string, defaultValue uint64) uint64 {
	var v uint64
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetFloat32 获取指定KEY的Float32值
func (c Config) GetFloat32(key string, defaultValue float32) float32 {
	var v float32
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

// GetFloat64 获取指定KEY的Float64值
func (c Config) GetFloat64(key string, defaultValue float64) float64 {
	var v float64
	if err := c.Decode(key, &v); err != nil {
		return defaultValue
	}
	return v
}

func find(keys []string, data Config) (*yaml.Node, error) {
	next, ok := data[keys[0]]
	if ok {
		// 叶子节点
		if 1 == len(keys) {
			return &next, nil
		}

		n := Config{}
		if err := next.Decode(&n); err != nil {
			return nil, err
		}

		return find(keys[1:], n)
	}

	return nil, errors.New("config not exist")
}
