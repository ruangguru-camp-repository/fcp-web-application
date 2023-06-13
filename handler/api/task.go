package api

import (
	"a21hc3NpZ25tZW50/apperror"
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid task ID"))
		return
	}
	var task model.Task
	err2 := c.ShouldBindJSON(&task)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	err3 := t.taskService.Update(id, &task)
	if err3 != nil {
		if errors.Is(err3, apperror.ErrInvalidUserIdOrCategoryId) {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(apperror.ErrInvalidUserIdOrCategoryId.Error()))
		return
	}
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	
	c.JSON(http.StatusOK, model.NewSuccessResponse("update task success"))

}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID"))
		return
	}
	err2 := t.taskService.Delete(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return

	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("delete task success"))
	// TODO: answer here
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	t2, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, t2)
	
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid task ID"))
		return
	}
	tc, err2 := t.taskService.GetTaskCategory(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}
	c.JSON(http.StatusOK, tc)

}
