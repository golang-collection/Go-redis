package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-04 18:20
* @Description: 读取远程配置文件
**/

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	err := viper.AddRemoteProvider("consul", CONSUL_URL, CONSUL_CONFIG)
	if err != nil {
		return err
	}
	viper.SetConfigType("json") // Need to explicitly set this to json
	err = viper.ReadRemoteConfig()
	if err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
