package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"memos/comm"
	"memos/models"
	"memos/service"
	"net/http"
)

func AddTaskController(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var reqBody models.MemosParam
	_ = json.Unmarshal(jsonData, &reqBody)

	content := reqBody.Content
	taskType := reqBody.Type
	predictTime := reqBody.PredictTime
	finishedTime := reqBody.FinishedTime

	res := service.AddTask(content, taskType, predictTime, finishedTime)
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
	comm.Logger.Info().Msg(fmt.Sprintf(
		"%s : Successfully!", res))
}
