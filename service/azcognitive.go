package service

import (
	"SearchGinBack/config"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
	"net/url"
)

func QueryResultFromAzCog(config *config.SearchYAML, keyword string) (string, error) {
	glog.Info("Searching keywords in Azure Cognitive Search ...")
	azCogUrl := fmt.Sprintf("%sindexes/%s/docs?api-version=2021-04-30-Preview&search=%s&queryLanguage=en-US&queryType=semantic&captions=extractive&semanticConfiguration=default",
		config.CogEndpoint,
		config.CogIndex,
		url.QueryEscape(keyword))

	req, err := http.NewRequest("GET", azCogUrl, nil)
	if err != nil {
		glog.Errorf("Error during construct Az Cog Get Request, message: %s", err)
		return "", err
	}

	req.Header.Set("api-key", config.CogSearchKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Error during sending Az Cog Get Request, message: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {

		if err != nil {
			glog.Errorf("Error decoding response: %s\n", err)
			return "", err
		}

		return string(respBytes), nil
	} else {
		return "", errors.New(fmt.Sprintf("Error during sending request to "))
	}

}
