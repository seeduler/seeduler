package utils

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
        Cors struct {
            AllowedOrigins []string `yaml:"allowed_origins"`
            AllowedMethods []string `yaml:"allowed_methods"`
            AllowedHeaders []string `yaml:"allowed_headers"`
        } `yaml:"cors"`
    } `yaml:"server"`
    JWT struct {
        SecretKey   string `yaml:"secret_key"`
        ExpiryHours int    `yaml:"expiry_hours"`
    } `yaml:"jwt"`
    Storage struct {
        EventFile string `yaml:"event_file"`
        HallFile  string `yaml:"hall_file"`
        UserFile  string `yaml:"user_file"`
    } `yaml:"storage"`
}

func LoadConfig(filename string) (*Config, error) {
    buf, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    config := &Config{}
    err = yaml.Unmarshal(buf, config)
    return config, err
}