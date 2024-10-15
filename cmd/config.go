package cmd

import (
	"os"

	"gopkg.in/yaml.v2"
)

const CONF_FILE = ".easy_label.yml"

type Label struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
}

type Config struct {
	LabelSets map[string][]Label `yaml:"label-sets"`
}

func getConfigFilename() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + CONF_FILE, nil
}

func getLabelSets() (map[string][]Label, error) {
	filename, err := getConfigFilename()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	return conf.LabelSets, nil
}
