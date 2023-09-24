package helper

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	ServicePort string `mapstructure:"SERVICE_PORT"`
	Host        string `mapstructure:"HOST"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	Port        string `mapstructure:"PORT"`
	SSLMode     string `mapstructure:"SSL_MODE"`
	Timezone    string `mapstructure:"TIMEZONE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
