package api

import (

	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid Category ID"))
		return
	}
	var category model.Category
	err2 := c.ShouldBindJSON(&category)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid request body"))
		return
	}

	err3 := ct.categoryService.Update(id, category)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("category update success"))
	
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid category ID"))
		return
	}

	err2 := ct.categoryService.Delete(id)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("category delete success"))
	
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	c2, err := ct.categoryService.GetList()
	if err != nil{
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusOK, c2)
	
}
