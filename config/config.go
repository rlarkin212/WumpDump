package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Version string  `mapstructure:"version"`
	Api     api     `mapstructure:"api"`
	Discord discord `mapstructure:"discord"`
}

type api struct {
	Port string `mapstructure:"port"`
}

type discord struct {
	GuildId string `mapstructure:"guildId"`
	Token   string `mapstructure:"token"`
}

func LoadConfig(path, name, extension string) (*Config, error) {
	c := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(extension)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	replaceEnvVars("${", "}")

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func replaceEnvVars(prefix, suffix string) {
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		if strings.HasPrefix(v, prefix) && strings.HasSuffix(v, suffix) {
			viper.Set(k, os.ExpandEnv(v))
		}
	}
}
