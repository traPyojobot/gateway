package loader

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const configPath = "../../config/models.yaml"

type model struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type group struct {
	GroupName string  `yaml:"name"`
	Model     []model `yaml:"model,inline"`
}

type Config struct {
	Group group `yaml:"group,inline"`
}

func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
