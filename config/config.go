package config

import "github.com/spf13/viper"

type Config struct {
    MySQL struct {
        DSN string `yaml:"dsn"`
    } `yaml:"mysql"`
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}