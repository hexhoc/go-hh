package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App        App
		HTTP       HTTP
		Log        Log
		Datasource Datasource
		Prometheus Prometheus
	}

	App struct {
		Name    string
		Version string
	}

	HTTP struct {
		Port    string
		Version string
	}

	Log struct {
		Level string
	}

	Datasource struct {
		Url     string
		PoolMax int
	}

	Prometheus struct {
		gateway string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigName("config.yml")  // name of config file
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")    // optionally look for config in the working directory
	viper.AddConfigPath("./../config") // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	cfg.App = App{
		viper.GetString("app.name"),
		viper.GetString("app.version"),
	}

	cfg.HTTP = HTTP{
		viper.GetString("http.port"),
		viper.GetString("http.version"),
	}

	cfg.Log = Log{
		viper.GetString("logger.log_level"),
	}

	cfg.Datasource = Datasource{
		viper.GetString("datasource.pg.url"),
		viper.GetInt("datasource.pg.pool_max"),
	}

	cfg.Prometheus = Prometheus{
		viper.GetString("prometheus.gateway"),
	}

	return cfg, nil
}
