package loader

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
