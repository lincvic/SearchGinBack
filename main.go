package main

import (
	"SearchGinBack/config"
	"SearchGinBack/model"
	"SearchGinBack/service"
	"github.com/avast/retry-go"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/glog"
	"net/http"
	"sync"
	"time"
)

func getDataFromAzCognitiveSearch(configYaml *config.SearchYAML, reqData *model.RequestData, wg *sync.WaitGroup, data *model.ServerRespData) {
	var err error
	var result string
	err = retry.Do(func() error {
		result, err = service.QueryResultFromAzCog(configYaml, reqData.Keyword)
		if err != nil {
			glog.Errorf("Error trying to query data from Az Cognitive Search, retrying ..., message: %s", err)
			return err
		}

		return nil
	}, retry.Attempts(3), retry.Delay(time.Second*1))

	if err != nil {
		glog.Errorf("Error during query data from Az Cognitive Search for 3 times, message: %s", err)
		data.SetCognitiveSearchRespWithLock("")
	}

	data.SetCognitiveSearchRespWithLock(result)
	wg.Done()
}

func getDataFromQdrant(configYaml *config.SearchYAML, reqData *model.RequestData, wg *sync.WaitGroup, data *model.ServerRespData) {
	var result string

	err := retry.Do(func() error {
		embeddingData, err := service.ConvertString2Vector(configYaml, reqData.Keyword)
		if err != nil {
			glog.Errorf("Error trying to converting data to embedding, retrying ..., message: %s", err)
			return err
		}

		result, err = service.QuerySingleResultFromQd(configYaml, embeddingData)
		if err != nil {
			glog.Errorf("Error trying to searching data in Qdrant, retrying ..., message: %s", err)
			return err
		}

		return nil
	}, retry.Attempts(3), retry.Delay(time.Second*1))

	if err != nil {
		glog.Errorf("Error during query data from Az Cognitive Search for 3 times, message: %s", err)
		if data != nil {
			data.SetQdrantRespWithLock("")
		}
	}

	data.SetQdrantRespWithLock(result)
	wg.Done()

}

func main() {

	configYaml, err := config.LoadYAML("config.yaml")
	if err != nil {
		glog.Errorf("Error during load yaml file, message: %s", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	router.POST("/calc_cog_emb", func(c *gin.Context) {
		reqData := new(model.RequestData)
		if err := c.BindJSON(reqData); err != nil {
			glog.Errorf("Error during handling request, message: %s", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		glog.Infof("Get key word %s, starting search ...", reqData.Keyword)

		wg := new(sync.WaitGroup)
		wg.Add(2)

		respData := &model.ServerRespData{}
		go getDataFromAzCognitiveSearch(configYaml, reqData, wg, respData)
		go getDataFromQdrant(configYaml, reqData, wg, respData)

		wg.Wait()

		glog.Info("Searching complete !")

		c.JSON(200, respData)
	})

	err = router.Run(":3339")
	if err != nil {
		glog.Errorf("Cannot start Gin server, message: %s", err)
		return
	}
}
