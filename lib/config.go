package lib

import (
	"io/ioutil"
	"os"

	"nnsay/ngcli/types"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ReadConfig -- load the config from the configuration file
func ReadConfig() (types.Config, error) {
	var config = types.Config{}
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
func SaveConfig(config types.Config) error {
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
