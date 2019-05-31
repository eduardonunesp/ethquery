package config

import "fmt"

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

func (cs ConfigurationList) GetCurrent(overwriteCurrent string) (*Configuration, error) {
	for _, configuration := range cs.Configurations {
		if overwriteCurrent != "" {
			if configuration.Name == overwriteCurrent {
				return &configuration, nil
			}
		} else {
			if configuration.Current {
				return &configuration, nil
			}
		}
	}

	return nil, fmt.Errorf("Current configuration not found, please command configuration")
}
