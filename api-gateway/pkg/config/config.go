package config

import "github.com/spf13/viper"

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	LogisticSvcUrl string `mapstructure:"LOGISTIC_SVC_URL"`
}

func LoadConfig() (c Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return

}
