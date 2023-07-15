package config

import "github.com/spf13/viper"

// 通过config进行代理
type config struct {
	viper *viper.Viper
}

var OcrConf *config
var OcrSecret *config

// 初始化配置
func init() {
	OcrConf = &config{
		viper: getConf(),
	}
	OcrSecret = &config{
		viper: getSecret(),
	}
}

// 读取配置文件
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
	v.SetConfigName("secret")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/secret")
	v.ReadInConfig()
	return v
}

// 通过config代理 拿到配置文件中定义字段值
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
