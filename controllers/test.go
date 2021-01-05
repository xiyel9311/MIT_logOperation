package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (tc *TestController) Test(c *gin.Context) {
	fmt.Println("test router.")
}
