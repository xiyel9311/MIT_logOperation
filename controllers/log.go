package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"logOperation/setting"

	"github.com/gin-gonic/gin"
)

var dcTr *http.Transport

var ip string

type LogController struct{}

type logOutput struct {
	UID   int `json:"UID"`
	MAC   int `json:"MAC"`
	Fltr  int `json:"Fltr"`
	TSt   int `json:"TSt"`
	TEnd  int `json:"TEnd"`
	Amt   int `json:"Amt"`
	Total int `json:"Total"`
	TLst  int `json:"TLst"`
	TFst  int `json:"TFst"`
}

func init() {
	ip = setting.IP()
}

func setReqHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic cm9vdDptaXRyb290")
	req.Header.Set("Cache-Control", "no-store")
}

func (lc *LogController) UpdateOutput(c *gin.Context) {

	data, _ := ioutil.ReadAll(c.Request.Body)

	fmt.Printf("data: %s\n", data)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, ip+"/log_output", bytes.NewBuffer(data))
	setReqHeader(req)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (lc *LogController) GetOutput(c *gin.Context) {

	req, err := http.NewRequest(http.MethodGet, ip+"/log_output", nil)
	setReqHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("field to connnect to get log output.")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	result := logOutput{}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return
	}
	fmt.Printf("log_output: %#v\n", result)
}

func (lc *LogController) LogMessage(c *gin.Context) {

	req, err := http.NewRequest(http.MethodGet, ip+"/log_message", nil)
	setReqHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("field to connnect to get log message.")
	}

	// defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("log_message: %s\n", bodyBytes)
}
