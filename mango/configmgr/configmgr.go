package configmgr

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Type string
	Path string
}

type ConfigMgr struct {
	Config
}

func (pkg ConfigMgr) Version() string {
	return "0.01"
}

func New(config Config) *ConfigMgr {
	if len(config.Name) == 0 {
		config.Name = "default"
	}

	if len(config.Type) == 0 {
		config.Type = "yaml"
	}

	if len(config.Path) == 0 {
		config.Path = "./config"
	}

	getAndSetConfig(config)

	pkg := ConfigMgr{
		Config: config,
	}

	return &pkg
}

func getAndSetConfig(config Config) {
	fmt.Println("CONFIG.NAME", config)
	viper.SetConfigName(config.Name) // name of config file (without extension)
	viper.SetConfigType(config.Type)
	viper.AddConfigPath(config.Path)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func (pkg ConfigMgr) Get(key string) interface{} {
	return viper.Get(key)
}

func (pkg ConfigMgr) GetString(key string) string {
	return viper.GetString(key)
}

func (pkg ConfigMgr) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (pkg ConfigMgr) GetInt(key string) int {
	return viper.GetInt(key)
}

func (pkg ConfigMgr) GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func (pkg ConfigMgr) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (pkg ConfigMgr) GetUint(key string) uint {
	return viper.GetUint(key)
}

func (pkg ConfigMgr) GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func (pkg ConfigMgr) GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func (pkg ConfigMgr) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (pkg ConfigMgr) GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func (pkg ConfigMgr) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func (pkg ConfigMgr) GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func (pkg ConfigMgr) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (pkg ConfigMgr) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func (pkg ConfigMgr) GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func (pkg ConfigMgr) GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

func (pkg ConfigMgr) GetSizeInBytes(key string) uint {
	return viper.GetSizeInBytes(key)
}
