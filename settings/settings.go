package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StarTime     string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"poolsize"`
}

func Init() (err error) {
	//viper.SetConfigFile("./config.yml")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v \n", err)
		return
	}
	//把读取到的配置信息反序列化到 Conf 变量中去
	err = viper.Unmarshal(Conf)
	if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了。。。")
		err := viper.Unmarshal(Conf)
		if err != nil {
			fmt.Printf("viper.Unmarshal(Conf) failed,err:%v\n", err)
			return
		}
	})
	return
}
