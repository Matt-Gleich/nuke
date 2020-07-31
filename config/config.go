package config

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

var path = "/.config/nuke/"

// Conf ... Config for nuke
type Conf struct {
	IgnoreUpdates bool     `yaml:"ignoreUpdates"`
	IgnoredApps   []string `yaml:"ignoredApps"`
}

// Exists ... If the config exists
func Exists() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get user home path", err, 0)
	}
	path = home + path

	yml, err := os.Stat(path + "config.yml")
	if !os.IsNotExist(err) {
		if !yml.IsDir() {
			path = path + "config.yml"
			return true
		}
	}
	yaml, err := os.Stat(path + "config.yaml")
	if !os.IsNotExist(err) {
		if !yaml.IsDir() {
			path = path + "config.yaml"
			return true
		}
	}
	return false
}

// Read ... Read from the config file
func Read(conf *Conf) *Conf {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		statuser.Error("Failed to read from config in\n\t"+path, err, 0)
	}

	err = yaml.Unmarshal(contents, &conf)
	if err != nil {
		statuser.Error("Failed to read the config", err, 0)
	}
	return conf
}
