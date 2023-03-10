package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"herrkung.com/GinVueEssential/model"
	"herrkung.com/GinVueEssential/repository"
	"herrkung.com/GinVueEssential/response"
	"herrkung.com/GinVueEssential/vo"
)

type ICategoryController interface {
	RestController
}

func NewCategoryController() ICategoryController {
	categoryRepository := repository.NewCategoryRepository()
	_ = categoryRepository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: categoryRepository}
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	//ShouldBind能够基于请求的不同，自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "create fail, category name is required")
		fmt.Println(requestCategory)
		return
	} else {
		category, err := c.Repository.Create(requestCategory.Name)
		if err != nil {
			response.Fail(ctx, nil, "create fail")
			return
		}
		response.Success(ctx, gin.H{"category": category}, "create successfully")
	}
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "update fail, category name is required")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "category is NOT found")
	} else {
		category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
		if err != nil {
			panic(err)
		}
		response.Success(ctx, gin.H{"category": category}, "update successfully")
	}
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "category is NOT found")
		return
	} else {
		response.Success(ctx, gin.H{"category": category}, "show category successfully")
	}
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "delete fail")
	} else {
		response.Success(ctx, nil, "delete successfully")
	}
}
