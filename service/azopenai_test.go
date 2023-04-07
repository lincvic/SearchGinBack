package service

import (
	"SearchGinBack/config"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestConvertString2Vector(t *testing.T) {
	testConfig, err := config.LoadYAML("../config/config.yaml")
	if err != nil {
		glog.Errorf("Error during load yaml file, message: %s", err)
	}

	ConvertString2Vector(testConfig, "This is Keyword")
}
