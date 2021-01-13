package controllers

import (
	"errors"
	"fmt"
	"net/url"

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

	params := "name=Dean&sex=male&hair_color=black+brown&eye_color="
	person, err := url.ParseQuery(params)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)

	// panic("a")
	return
}
