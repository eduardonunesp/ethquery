package config

type Configuration struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Current bool   `yaml:"current"`
}

type ConfigurationList struct {
	Configurations []Configuration `yaml:"configurations"`
}

func NewConfiguration(name, url string) *Configuration {
	return &Configuration{name, url, false}
}

func NewConfigurations(configurations []Configuration) *ConfigurationList {
	return &ConfigurationList{configurations}
}

func (cs ConfigurationList) GetCurrent() Configuration {
	for _, configuration := range cs.Configurations {
		if configuration.Current {
			return configuration
		}
	}

	panic("Current config not found")
}
