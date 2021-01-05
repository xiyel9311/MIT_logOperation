package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (lc *LogController) UpdateOutput(c *gin.Context) {

	UID, _ := strconv.Atoi(c.Query("UID"))
	MAC, _ := strconv.Atoi(c.Query("MAC"))
	Fltr, _ := strconv.Atoi(c.Query("Fltr"))
	TSt, _ := strconv.Atoi(c.Query("TSt"))
	TEnd, _ := strconv.Atoi(c.Query("TEnd"))
	Amt, _ := strconv.Atoi(c.Query("Amt"))
	Total, _ := strconv.Atoi(c.Query("Total"))
	TLst, _ := strconv.Atoi(c.Query("TLst"))
	TFst, _ := strconv.Atoi(c.Query("TFst"))

	input := logOutput{
		UID:   UID,
		MAC:   MAC,
		Fltr:  Fltr,
		TSt:   TSt,
		TEnd:  TEnd,
		Amt:   Amt,
		Total: Total,
		TLst:  TLst,
		TFst:  TFst,
	}
	fmt.Printf("input: %#v\n", input)

	jsonStrBytes, err := json.Marshal(input)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, "http://192.168.10.145:80/log_output", bytes.NewBuffer(jsonStrBytes))
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (lc *LogController) GetOutput(c *gin.Context) {

	response, err := http.Get("http://192.168.10.145:80/log_output")

	// defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)

	result := logOutput{}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return
	}
	fmt.Printf("log_output: %#v\n", result)
}
