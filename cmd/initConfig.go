package cmd

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

// InitConfig 传入配置文件路径
func InitConfig(cfg string) error {
	//fsnotify.NewWatcher()

	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.readConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

// 加载配置
func (c *Config) readConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("config") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("settings.dev")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	//viper.AutomaticEnv()            // 读取匹配的环境变量
	//viper.SetEnvPrefix("tapi-blog") // 设置环境变量的前缀为tapi-blog
	//replacer := strings.NewReplacer(".", "_")
	//viper.SetEnvKeyReplacer(replacer)
	//viper.SetDefault("redis.password", "") 设置默认值
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: %s", e.Name)
	})
}
