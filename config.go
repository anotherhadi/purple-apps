package purple

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Global struct {
		Colors struct {
			Accent    string `yaml:"accent"`
			Gray      string `yaml:"gray"`
			LightGray string `yaml:"lightgray"`
			Muted     string `yaml:"muted"`
		}
	}
	Calendar struct {
		Paths        []string `yaml:"paths"`
		DefaultView  string   `yaml:"default_view"`
		SidebarWidth int      `yaml:"sidebar_width"`
	}
}

func getDefaultConfig() Config {
	data := `
global:
  colors:
    accent: "#A594FD"
    gray: "245"
    lightgray: "241"
    muted: "#202022"
calendar:
  paths:
    - "$HOME/calendar"
  default_view: "month"
  sidebar_width: 25
  `

	// Create a struct to hold the YAML data
	var config Config

	// Unmarshal the YAML data into the struct
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	return config
}

func getUserConfig() Config {

	config_home := os.Getenv("XDG_CONFIG_HOME")
	if config_home == "" {
		config_home = os.Getenv("HOME") + "/.config"
	}
	config_home = strings.TrimSuffix(config_home, "/")

	// Read the file
	data, err := os.ReadFile(config_home + "/purple.yaml")
	if err != nil {
		return Config{}
	}

	// Create a struct to hold the YAML data
	var config Config

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}
	}
	return config
}

// the second config will override the first one
func mergeConfig(first, second Config) Config {
	if second.Global.Colors.Accent != "" {
		first.Global.Colors.Accent = second.Global.Colors.Accent
	}
	if second.Global.Colors.Gray != "" {
		first.Global.Colors.Gray = second.Global.Colors.Gray
	}
	if second.Global.Colors.LightGray != "" {
		first.Global.Colors.LightGray = second.Global.Colors.LightGray
	}
	if second.Global.Colors.Muted != "" {
		first.Global.Colors.Muted = second.Global.Colors.Muted
	}
	return first
}

func GetConfig() Config {
	userConfig := getUserConfig()
	defaultConfig := getDefaultConfig()
	mergedConfig := mergeConfig(defaultConfig, userConfig)
	return mergedConfig
}
