package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("conf changed: %v\n", in.Op)
	})
	return nil
}
