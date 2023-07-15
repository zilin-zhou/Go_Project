package config

import "github.com/spf13/viper"

type config struct {
	viper *viper.Viper
}

var (
	SesConf   *config
	SesSecret *config
)

func init() {
	SesConf = &config{
		viper: getSmsConf(),
	}
	SesSecret = &config{
		viper: getSmsSecret(),
	}
}

func getSmsConf() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/config")
	v.ReadInConfig()
	return v
}

func getSmsSecret() *viper.Viper {
	v := viper.New()
	v.SetConfigName("secret")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/secret")
	v.ReadInConfig()
	return v
}
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
