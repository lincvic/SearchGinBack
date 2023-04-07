package service

import (
	"SearchGinBack/config"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestQuerySingleResultFromQd(t *testing.T) {
	testConfig, err := config.LoadYAML("../config/config.yaml")
	if err != nil {
		glog.Errorf("Error during load yaml file, message: %s", err)
	}

	vector, _ := ConvertString2Vector(testConfig, "茄子面")
	QuerySingleResultFromQd(testConfig, vector)
}
