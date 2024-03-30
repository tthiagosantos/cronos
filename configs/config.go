package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var ConfigEnv *Config

type Config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	GrpcServerPort string `mapstructure:"GRPC_PORT"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JwtExperesIn   int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth      *jwtauth.JWTAuth
	BASE_URL       string `mapstructure:"BASE_URL"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&ConfigEnv)
	if err != nil {
		panic(err)
	}
	ConfigEnv.TokenAuth = jwtauth.New("HS256", []byte(ConfigEnv.JWTSecret), nil)
	return ConfigEnv, err
}
