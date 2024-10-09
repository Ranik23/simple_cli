package config

import (
	"io"
	"os"
	"gopkg.in/yaml.v2"
)


func LoadConfig(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Config struct {
	Env 		string 		`yaml:"env"`
	Logging 	Logging 	`yaml:"logging"`
	PprofServer PprofServer `yaml:"PprofServer"`
	DataBase	DataBase    `yaml:"database"`
}

type Logging struct {
	Level string `yaml:"level"`
}

type PprofServer struct {
	Host string `yaml:"host"`
}

type DataBase struct {
	Host 	 string `yaml:"host"`
	Port 	 int    `yaml:"port"`
	User 	 string `yaml:"user"`
	Password string `yaml:"password"`
	DBName	 string `yaml:"dbname"`
}