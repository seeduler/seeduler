package utils

import (
    "encoding/json"
    "os"
)

type Config struct {
    JWTSecretKey string `json:"jwt_secret_key"`
}

func LoadConfig(filePath string) (*Config, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := &Config{}
    err = decoder.Decode(config)
    if err != nil {
        return nil, err
    }

    return config, nil
}