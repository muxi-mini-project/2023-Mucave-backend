package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string, prefix string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(prefix); err != nil {
		return err
	}

	return nil
}

func (c *Config) initConfig(prefix string) error {
	if c.Name != "" {
		// 如果制定配置文件，则解析配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 	如果没有制指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("C:/Users/onlyB/GolandProjects/a/2023-Mucave-backend/conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // 默认配置文件类型为yaml
	viper.AutomaticEnv()        // 读取默认的环境变量
	viper.SetEnvPrefix(prefix)  // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
