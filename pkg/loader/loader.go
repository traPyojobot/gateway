package loader

import (
	"io/ioutil"

	"github.com/traPyojobot/gateway/pkg/model"
	"gopkg.in/yaml.v3"
)

const ConfigPath = "./config/models.yaml"

type mlmodel struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type group struct {
	GroupName string    `yaml:"name"`
	Model     []mlmodel `yaml:"models"`
}

type Config struct {
	Group []group `yaml:"group"`
}

func LoadMLModel() (map[string][]model.MLModel, error) {
	res := map[string][]model.MLModel{}
	c, err := loadConfig()
	if err != nil {
		return nil, err
	}
	for _, g := range c.Group {
		var s []model.MLModel
		for _, m := range g.Model {
			if err != nil {
				return nil, err
			}
			s = append(s, model.MLModel{
				Name: m.Name,
				Url:  m.Url,
			})
		}
		res[g.GroupName] = s
	}
	return res, nil
}

func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(ConfigPath)
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
