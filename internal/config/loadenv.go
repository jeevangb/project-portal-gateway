package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthPORT       string `mapstructure:"GRPC_SERVER_PORT"`
	PrivateKeyPath string `mapstructure:"PRIVATE_KEY_PATH"`
	PublicKeyPath  string `mapstructure:"PUBLIC_KEY_PATH"`
}

func LoadConfig(env *string) (Config, error) {
	v := viper.New()
	var cfg Config
	envs := []string{
		"PORT",
		"GRPC_SERVER_PORT",
		"PRIVATE_KEY_PATH",
		"PUBLIC_KEY_PATH",
	}
	v.AddConfigPath("./")
	v.SetConfigFile("internal/config/" + *env + "/application.env")
	v.ReadInConfig()
	for _, val := range envs {
		if err := v.BindEnv(val); err != nil {
			return cfg, err
		}
	}
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("error while unmarshalling data to struct")
		return cfg, err
	}
	if err := validator.New().Struct(&cfg); err != nil {
		log.Fatalf("error while unmarshalling data to struct")
		return cfg, err
	}
	return cfg, nil
}
