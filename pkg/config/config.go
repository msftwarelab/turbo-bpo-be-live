package config

import (
	"encoding/json"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/spf13/viper"
)

var AppConfig *viper.Viper

func Init(serviceType string, env string) {

	log.Info("%s", "config initialization...")
	defer log.Info("%s", "config initialization...done!")
	AppConfig = viper.New()
	AppConfig.AddConfigPath(".")
	AppConfig.SetConfigName("staging.config")
	if env == "production" {
		AppConfig.SetConfigName("production.config")
		if serviceType == "jobs" {
			AppConfig.SetConfigName("jobs.production.config")
		}
		if serviceType == "pinger" {
			AppConfig.SetConfigName("pinger.production.config")
		}

	}
	if env == "staging" {
		AppConfig.SetConfigName("staging.config")
		if serviceType == "jobs" {
			AppConfig.SetConfigName("jobs.staging.config")
		}
		if serviceType == "pinger" {
			AppConfig.SetConfigName("pinger.staging.config")
		}
	}
	if env == "development" {
		AppConfig.SetConfigName("development.config")
		if serviceType == "jobs" {
			AppConfig.SetConfigName("jobs.development.config")
		}
		if serviceType == "pinger" {
			AppConfig.SetConfigName("pinger.development.config")
		}

	}
	AppConfig.SetConfigType("json")
	err := AppConfig.ReadInConfig()
	if err != nil {
		log.Error("viper error: %v", err)
	}

	log.Info("config: %s", PrettyPrint(AppConfig.AllSettings()))
	return
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
