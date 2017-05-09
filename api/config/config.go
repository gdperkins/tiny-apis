package config

import (
	"log"

	"github.com/spf13/viper"
)

var c *viper.Viper

// Init loads the correct environment yaml
// into the viper configuration package
func Init(env string) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("config/")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	c = v
}

// GetCurrent gets the current app envionment config
func GetCurrent() *viper.Viper {
	return c
}
