package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

func Load() ConfigurationList {
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("fatal error on obtain home dir: %s", err))
	}

	homeDirConfig := fmt.Sprintf("%s/.ethquery", usr.HomeDir)
	homeConfigFile := fmt.Sprintf("%s/%s", homeDirConfig, configFilename)

	if isDir := dirExists(homeDirConfig); !isDir {
		if err := os.Mkdir(homeDirConfig, 0755); err != nil {
			panic(fmt.Errorf("fatal error on create config file: %s", err))
		}
	}

	if configExists := fileExists(homeConfigFile); !configExists {
		if err := ioutil.WriteFile(homeConfigFile, []byte{}, 0644); err != nil {
			panic(fmt.Errorf("unable to write config file: %s", err))
		}
	}

	bs, err := ioutil.ReadFile(homeConfigFile)
	if err != nil {
		panic(fmt.Errorf("unable to read config file: %s", err))
	}

	var configList ConfigurationList
	if err := yaml.Unmarshal(bs, &configList); err != nil {
		panic(fmt.Errorf("unable to unmarshall config file: %s", err))
	}

	return configList
}
