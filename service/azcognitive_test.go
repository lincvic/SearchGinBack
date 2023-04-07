package service

import (
	"SearchGinBack/config"
	"github.com/gogf/gf/os/glog"
	"testing"
)

func TestQueryResultFromAzCog(t *testing.T) {
	testConfig, err := config.LoadYAML("../config/config.yaml")
	if err != nil {
		glog.Errorf("Error during load yaml file, message: %s", err)
	}

	QueryResultFromAzCog(testConfig, "Administrator account")
}
