package lib

import (
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type AuthConfig struct {
	Token   string `yaml:"token,omitempty"`
	TokenAt string `yaml:"tokenAt,omitempty"`
	OrgId   int    `yaml:"orgId"`
}
type Config struct {
	Endpoint        string     `yaml:"endpoint"`
	Username        string     `yaml:"username"`
	Password        string     `yaml:"password"`
	ApplicationType int        `yaml:"applicationType"`
	Auth            AuthConfig `yaml:"auth,omitempty"`
}

// ReadConfig -- load the config from the configuration file
func ReadConfig() (Config, error) {
	var config = Config{}
	filePath := viper.ConfigFileUsed()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadFile(filePath)
	err = yaml.Unmarshal(bytes, &config)
	return config, err
}

// SaveConfig -- save the config into configuration file
func SaveConfig(config Config) error {
	filePath := viper.ConfigFileUsed()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, _ := yaml.Marshal(config)
	_, err = file.Write(bytes)
	return err
}
