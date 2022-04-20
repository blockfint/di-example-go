package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Initialize() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// Replace double underscore (__) in os env with the viper dot notation (.)
	// e.g., os env HTTP_SERVER_PORT will be converted to HTTP_SERVER.PORT
	// So we can set the os env flag to override the config.yaml
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Unable to read config: %v\n", err))
	}
}
