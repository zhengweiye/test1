package test1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhengweiye/test2/test2"
)

type TestController interface {
	SayHi(ctx *gin.Context)
}

type TestControllerImpl struct {
}

func NewTestController() TestController {
	return TestControllerImpl{}
}

// SayHi
// @Tags 用户管理
// @Summary sayHi接口
// @Description 描述
// @Router /enforce_admin/user/sayHi [post]
// @Security ApiKeyAuth
// @Param data body Test1 true "请求参数"
// @Success 200 {object} test2.Test2
func (t TestControllerImpl) SayHi(ctx *gin.Context) {
	reqBean := GetBean[Test1](ctx)
	resBean := test2.Test2{
		Id:   reqBean.Id,
		Name: reqBean.Name,
	}
	ctx.JSON(200, resBean)
}

func GetBean[T any](ctx *gin.Context) T {
	var data T
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}
	return data
}
