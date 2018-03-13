package config

import (
	"strings"
	"fmt"
	"time"
	"github.com/spf13/viper"
	"log"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

func Config() Provider {
	return defaultConfig
}

func init() {
	defaultConfig = readViperConfig("go-lorem-image")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// app
	v.SetDefault("app.name", "goloremimage")
	// log
	v.SetDefault("log.json", false)
	v.SetDefault("log.level", "info")
	// http defaults
	v.SetDefault("api.listen", ":9999")
	v.SetDefault("api.domain", "localhost:9999")
	v.SetDefault("api.corsOrigins", []string{"*"})

	v.SetDefault("images.root", "./assets")

	// read configuration from file
	v.SetConfigName("config")
	v.AddConfigPath("/etc/go-lorem-image/")
	v.AddConfigPath("/opt/go-lorem-image/etc/")
	v.AddConfigPath("./config/")
	err := v.ReadInConfig()
	if err != nil {
		log.Print(fmt.Errorf("Configuration file not read: %s \n", err))
	}

	return v
}
