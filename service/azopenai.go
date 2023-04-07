package service

import (
	"SearchGinBack/config"
	"SearchGinBack/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"net/http"
)

func ConvertString2Vector(config *config.SearchYAML, keywords string) ([]float64, error) {
	glog.Info("Converting keywords into embedding ...")
	azOpenAIEndpoint := fmt.Sprintf("%sopenai/deployments/%s/embeddings?api-version=2022-12-01", config.OaiEndpoint, config.OaiDeploymentName)
	query := map[string]interface{}{
		"input": keywords,
	}

	jsonPayload, err := json.Marshal(query)
	if err != nil {
		glog.Errorf("Error during marshal keywords into json bytes, message: %s", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", azOpenAIEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		glog.Errorf("Error during making a new post request, message: %s", err)
		return nil, err
	}

	req.Header.Set("api-key", config.OaiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Azure openai cannot process current string into embedding, message: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		result := new(model.OAIData)
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			glog.Errorf("Error decoding response: %s\n", err)
			return nil, err
		}

		return result.Data[0].Embedding, nil

	} else {
		glog.Errorf("Azure openai cannot process current string into embedding, code: %d", resp.StatusCode)
		return nil, errors.New(fmt.Sprintf("Azure openai cannot process current string into embedding, code: %d", resp.StatusCode))
	}

}
