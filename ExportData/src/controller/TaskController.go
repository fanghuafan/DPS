package controller

import (
	"ExportData/src/common/logger"
	"ExportData/src/models"
	"ExportData/src/service"
	"ExportData/src/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type TaskController struct {
	Log         logger.ILogger       `inject:""`
	TaskService service.ITaskService `inject:""`
}

// get the task info
func (taskController *TaskController) Get(c *gin.Context) {
	taskCode := c.Query("taskCode")
	if len(taskCode) == 0 {
		c.JSON(100, gin.H{
			"status":  "error",
			"message": "query param null",
		})
		return
	}
	task := taskController.TaskService.GetTask(taskCode)
	c.JSON(200, gin.H{
		"status":  "geted",
		"message": "query ok",
		"task":    task,
	})
}

// add task record
func (taskController *TaskController) Add(c *gin.Context) {
	description := c.Query("desc")

	nowTime := time.Now()
	nowTimeStr := nowTime.Format("YYYY-MM-DD")

	var build strings.Builder
	build.WriteString(nowTimeStr)
	build.WriteString("-")
	build.WriteString(utils.RandString(3))
	taskCode := build.String()

	timeout, _ := strconv.ParseInt(c.Query("timeout"), 10, 32)
	creator, _ := strconv.ParseInt(c.Query("user"), 10, 32)

	task := &models.Task{
		Creator:     int32(creator),
		Description: description,
		TaskCode:    taskCode,
		TaskName:    c.Query("name"),
		CreateTime:  nowTime,
		State:       0,
		Timeout:     int32(timeout),
	}

	taskController.TaskService.Add(task)

	c.JSON(200, gin.H{
		"status":  "added",
		"message": "add ok",
		"task":    task,
	})
}
