package config

import "github.com/spf13/viper"

type AppConfig struct{
	APIPort string `mapstructure:"API_PORT"`
	RootStartWarsApi string `mapstructure:"ROOT_START_WARS_API"`
	AppDBHost string `mapstructure:"APP_DB_HOST"`
	AppDBName string `mapstructure:"APP_DB_NAME"`
	AppDBPort string `mapstructure:"APP_DB_PORT"`
	AppDBUserName string `mapstructure:"APP_DB_USER_NAME"`
	AppDBPassword string `mapstructure:"APP_DB_Password"`
}

func LoadConfig(path string) (config *AppConfig, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}