package configs

import (
	"github.com/convee/go-blog-api/pkg/jwt"
	"github.com/convee/go-blog-api/pkg/logger"
	"github.com/convee/go-blog-api/pkg/redis"
	"github.com/convee/go-blog-api/pkg/storage/orm"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config global config
type Config struct {
	// common
	App  AppConfig
	Cron CronConfig
	// component config
	Logger logger.Config
	ORM    orm.Config
	Redis  redis.Config
	JWT    jwt.Config
}

// AppConfig app config
type AppConfig struct {
	Name          string
	Version       string
	Mode          string
	Addr          string
	Host          string
	Resource      string
	FfprobePath   string
	Env           string
	MpHost        string
	AdOctopusHost string
}

// CronConfig cron config
type CronConfig struct {
	Push bool
}

var (
	// Conf app global config
	Conf = &Config{}
)

func Init(configPath string) *Config {
	viper.SetConfigType("yml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	})
	return Conf
}
