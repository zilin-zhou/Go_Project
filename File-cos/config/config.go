package config

import "github.com/spf13/viper"

type config struct {
	viper *viper.Viper
}

var Confs *config
var Secret *config

func init() {
	Confs = &config{
		getConf(),
	}
	Secret = &config{
		getSecret(),
	}
}
func getConf() *viper.Viper {
	v := viper.New()
	v.SetConfigName("confs")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/conf")
	v.ReadInConfig()
	return v
}
func getSecret() *viper.Viper {
	v := viper.New()
	v.SetConfigName("secrets")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/secret")
	v.ReadInConfig()
	return v
}
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
