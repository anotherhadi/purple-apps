package purple

import (
	"fmt"
	"os"
	"strings"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

type config struct {
	Global struct {
		Colors struct {
			Accent    string `yaml:"accent"`
			Gray      string `yaml:"gray"`
			LightGray string `yaml:"lightgray"`
			Muted     string `yaml:"muted"`
		}
	}
	Calendar struct {
		Paths       []string `yaml:"paths"`
		DefaultView string   `yaml:"default_view"`
	}
}

var (
	defaultConfig = `
global:
  colors:
    accent: "#A594FD"
    gray: "245"
    lightgray: "241"
    muted: "#202022"
calendar:
  paths:
    - "$HOME/calendar.md"
  default_view: "month"
  `
	Config = getConfig()
)

func getDefaultConfig() config {
	// Create a struct to hold the YAML data
	var config config

	// Unmarshal the YAML data into the struct
	err := yaml.Unmarshal([]byte(defaultConfig), &config)
	if err != nil {
		fmt.Println(err)
		return config
	}
	return config
}

func getUserConfig() config {
	config_home := os.Getenv("XDG_CONFIG_HOME")
	if config_home == "" {
		config_home = os.Getenv("HOME") + "/.config"
	}
	config_home = strings.TrimSuffix(config_home, "/")

	// Read the file
	data, err := os.ReadFile(config_home + "/purple.yaml")
	if err != nil {
		return config{}
	}

	// Create a struct to hold the YAML data
	var userconfig config

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &userconfig)
	if err != nil {
		return config{}
	}
	return userconfig
}

func getConfig() config {
	userConfig := getUserConfig()
	defaultConfig := getDefaultConfig()
	if err := mergo.Merge(&defaultConfig, userConfig, mergo.WithOverride); err != nil {
		return defaultConfig
	}
	return defaultConfig
}
