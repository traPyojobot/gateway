package loader

import (
	"io/ioutil"
	"net/url"

	"github.com/traPyojobot/gateway/pkg/proxy"
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
	Group []group `yaml:"group,inline"`
}

func LoadMLModel() (map[string][]proxy.MLModel, error) {
	var res map[string][]proxy.MLModel
	c, err := loadConfig()
	if err != nil {
		return nil, err
	}
	for _, g := range c.Group {
		var s []proxy.MLModel
		for _, m := range g.Model {
			u, err := url.Parse(m.Url)
			if err != nil {
				return nil, err
			}
			s = append(s, proxy.MLModel{
				Name: m.Name,
				Url:  u,
			})
		}
		res[g.GroupName] = s
	}
	return res, nil

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
