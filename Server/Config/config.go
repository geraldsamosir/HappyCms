package Config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigService struct {
}

func (c *ConfigService) ServiceConf() {
	// configuration file settings
	viper.AddConfigPath("configs/.")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Config file not found...")
	}

	// log service settings
	var logger Log
	logger.Logging()
}
