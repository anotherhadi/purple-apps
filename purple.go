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
	Pomodoro struct {
		Path           string `yaml:"path"`
		StartView      string `yaml:"start_view"`
		Amount         int    `yaml:"amount"`
		WorkTime       int    `yaml:"work_time"`
		ShortPauseTime int    `yaml:"short_pause_time"`
		LongPauseTime  int    `yaml:"long_pause_time"`
	}
}

var (
	defaultConfig = `
global:
  colors:
    accent: "#A594FD"
    gray: "#2A2A2C"
    lightgray: "#3A3A3C"
    muted: "#202022"
calendar:
  paths:
    - "$HOME/calendar.md"
  default_view: "month"
pomodoro:
  path: "/tmp/pomodoro"
  start_view: "start"
  amount: 4
  work_time: 25
  short_pause_time: 5
  long_pause_time: 15
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
