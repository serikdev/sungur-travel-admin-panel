package config

import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig
	Database DBConfig
}

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	Name     string
	SSLmode  string
}

func LoadCfg() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Setting default value
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSL_MODE", "disabled")

	cfg := &Config{
		Server: ServerConfig{
			Port: viper.GetString("SERVER_PORT"),
		},
		Database: DBConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			Password: viper.GetString("DB_PASSWORD"),
			User:     viper.GetString("DB_USER"),
			Name:     viper.GetString("DB_NAME"),
			SSLmode:  viper.GetString("DB_SSL_MODE"),
		},
	}
	return cfg, nil
}
