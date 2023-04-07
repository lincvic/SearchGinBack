package service

import (
	"SearchGinBack/config"
	"SearchGinBack/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
)

func QuerySingleResultFromQd(config *config.SearchYAML, vector []float64) (string, error) {
	glog.Infof("Searching embedding in Qdrant ...")
	qdEndpoint := fmt.Sprintf("http://%s:%d/collections/%s/points/search", config.QdrantAddress, config.QdrantPort, config.QdrantCollection)

	query := model.QueryPayload{
		TopK:        3,
		Vector:      vector,
		Space:       "cosine",
		WithPayLoad: true,
	}
	payload, err := json.Marshal(query)
	if err != nil {
		glog.Errorf("Error encoding query: %s\n", err)
		return "", err
	}

	req, err := http.NewRequest("POST", qdEndpoint, bytes.NewReader(payload))
	if err != nil {
		glog.Errorf("Error creating request: %s\n", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Error sending request: %s\n", err)
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
		glog.Errorf("Error response: %s\n", resp.Status)
		return "", errors.New(fmt.Sprintf("http error code %d", resp.StatusCode))
	}
}
