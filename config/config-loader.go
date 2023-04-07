package config

import (
	"github.com/gogf/gf/os/glog"
	yamlLoader "gopkg.in/yaml.v3"
	"io/ioutil"
)

type SearchYAML struct {
	QdrantAddress    string `yaml:"QdrantAddress"`
	QdrantPort       int    `yaml:"QdrantPort"`
	QdrantCollection string `yaml:"QdrantCollection"`

	CogSearchKey string `yaml:"CogSearchKey"`
	CogEndpoint  string `yaml:"CogEndpoint"`
	CogIndex     string `yaml:"CogIndex"`

	OaiKey            string `yaml:"OaiKey"`
	OaiEndpoint       string `yaml:"OaiEndpoint"`
	OaiDeploymentName string `yaml:"OaiDeploymentName"`
}

func LoadYAML(yamlFileName string) (*SearchYAML, error) {
	glog.Infof("Loading YAML file %s", yamlFileName)

	yamlFile, err := ioutil.ReadFile(yamlFileName)
	if err != nil {
		glog.Errorf("Failed Loading YAML file %s, Error is:%s", yamlFileName, err)
		return nil, err
	}

	searchConfig := new(SearchYAML)
	err = yamlLoader.Unmarshal(yamlFile, searchConfig)
	if err != nil {
		glog.Errorf("Failed Unmarshal YAML file %s, Error is:%s", yamlFileName, err)
		return nil, err
	}

	glog.Infof("Successfully Loaded YAML file %s", yamlFileName)
	return searchConfig, err
}
