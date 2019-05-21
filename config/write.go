package config

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"gopkg.in/yaml.v2"
)

func Write(configurationList *ConfigurationList) {
	bs, err := yaml.Marshal(configurationList)
	if err != nil {
		panic(fmt.Errorf("fatal error write configuration: %s", err))
	}

	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("fatal error on obtain home dir: %s", err))
	}

	homeDirConfig := fmt.Sprintf("%s/.ethquery", usr.HomeDir)
	homeConfigFile := fmt.Sprintf("%s/%s", homeDirConfig, configFilename)

	if err := ioutil.WriteFile(homeConfigFile, bs, 0644); err != nil {
		panic(fmt.Errorf("fatal error write configuration: %s", err))
	}
}
