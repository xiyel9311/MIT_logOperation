package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (tc *TestController) Test(c *gin.Context) {
	fmt.Println("test router.")
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic!")
			}
		} else {
			fmt.Println("commit.")
		}
	}()

	// panic("a")
	return
}
