package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Version string `yaml:"version"`
	Mode    string `yaml:"mode"`
	Server  struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
	Eth struct {
		Url struct {
			Http string `yaml:"http"`
			Ws   string `yaml:"ws"`
		} `yaml:"url"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"eth"`
}

func Load(file string) (*Config, error) {

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("confile error %q: %v", file, err)
	}
	// fmt.Println("conf", c)
	return c, nil
}
