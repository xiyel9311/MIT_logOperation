package rs485DataControllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logOperation/models/sysparam"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RS485Controller struct{}

var ip string

func init() {
	ip = sysparam.MainIP
}

func setReqHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic cm9vdDptaXRyb290")
	req.Header.Set("Cache-Control", "no-store")
}

type comData struct {
	ExpWord []expWordContent
}

type expWordContent struct {
	Ch    int `json:"Ch"`
	Val   int `json:"Val"`
	Evt   int `json:"Evt"`
	SID   int `json:"SID"`
	Addr  int `json:"Addr"`
	MAddr int `json:"MAddr"`
	WEvt  int `json:"WEvt"`
}

func (rc *RS485Controller) ComData(c *gin.Context) {
	comNum := c.Param("comID")

	req, err := http.NewRequest(http.MethodGet, ip+"/expansion_word/com_"+comNum, nil)
	setReqHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to connnect to get expansion word.")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	result := comData{}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return
	}
	fmt.Printf("com_%s Data: %+v\n", comNum, result)

	c.JSON(http.StatusOK, result)
}

func (rc *RS485Controller) ChData(c *gin.Context) {
	comNum := c.Param("comID")
	chNum := c.Param("chID")

	req, err := http.NewRequest(http.MethodGet, ip+"/expansion_word/com_"+comNum+"/ch_"+chNum, nil)
	setReqHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to connnect to get expansion word.")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	result := expWordContent{}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return
	}
	fmt.Printf("com_%s ch_%s Data: %+v\n", comNum, chNum, result)

}

func (rc *RS485Controller) TempratureAndHumidity(c *gin.Context) {

	req, err := http.NewRequest(http.MethodGet, ip+"/expansion_word/com_"+strconv.Itoa(sysparam.SensorComNum), nil)
	setReqHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to connnect to get temprature & humidity.")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	result := comData{}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return
	}
	fmt.Printf("com_%d temprature: %3.2f degrees Celsius, humidity: %2.2f%%\n", sysparam.SensorComNum, float64(result.ExpWord[sysparam.TempratureChNum].Val)/10, float64(result.ExpWord[sysparam.HumidityChNum].Val)/10)
}
