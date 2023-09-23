package helper

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Host        string `mapstructure:"HOST"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	Port        string `mapstructure:"PORT"`
	SSLMode     string `mapstructure:"SSL_MODE"`
	Timezone    string `mapstructure:"TIMEZONE"`
	ServicePort string `mapstructure:"SERVICE_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
