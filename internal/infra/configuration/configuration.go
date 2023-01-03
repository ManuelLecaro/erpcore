package configuration

import "github.com/spf13/viper"

const (
	portConfig             = "PORT"
	portConfigDefaultValue = "3000"
	NameConfig             = "NAME"
	NameConfigDefaultValue = "erpcore"
)

// Section to handle global conf on infrastructure
func LoadConfiguration() *viper.Viper {
	configs := viper.New()
	configs.SetConfigFile(".env")
	configs.ReadInConfig()

	configs.SetDefault(portConfig, portConfigDefaultValue)
	configs.SetDefault(NameConfig, NameConfigDefaultValue)

	return configs
}
