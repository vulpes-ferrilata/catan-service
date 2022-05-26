package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type CatanController interface {
	Get(ctx iris.Context) (mvc.Result, error)
}

func NewCatanController() CatanController {
	return &catanController{}
}

type catanController struct{}

func (c catanController) Get(ctx iris.Context) (mvc.Result, error) {
	return mvc.Response{
		Code: iris.StatusNotImplemented,
	}, nil
}
