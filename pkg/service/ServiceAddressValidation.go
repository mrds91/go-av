package service

import (
	"address-validation/pkg/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CallAddressValidationService(c *gin.Context, expAddressReq model.ExpAddressReq) {

	// var expAddressReq model.ExpAddressReq
	// if err := c.BindJSON(&expAddressReq); err != nil {
	// 	//return
	// }

	var sysAddressInfo model.SysAddressInfo

	sysAddressInfo.Addressline1 = expAddressReq.Addresses[0].Address.Addressline1
	sysAddressInfo.Addressline2 = expAddressReq.Addresses[0].Address.Addressline2
	sysAddressInfo.Cityname = expAddressReq.Addresses[0].Address.City
	sysAddressInfo.Statecode = expAddressReq.Addresses[0].Address.State
	sysAddressInfo.Postalcode = expAddressReq.Addresses[0].Address.ZipCode
	pbuff := new(bytes.Buffer)
	json.NewEncoder(pbuff).Encode(sysAddressInfo)

	url := "http://localhost:9000/api/common/v1/addresses"

	req, err := http.NewRequest("POST", url, pbuff)

	if err != nil {
		fmt.Println(err.Error())
	}
	//Headers
	req.Header.Add("x-rapidapi-key", "YOU_API_KEY")
	req.Header.Add("Content-Type", "application/json")
	//timeout settings
	var myClient = &http.Client{Timeout: 10 * time.Second}
	res, err := myClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//parse the response
	defer res.Body.Close()
	body, readerr := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if readerr != nil {
		fmt.Println(readerr.Error())
		fmt.Println(body)
		//return
	}

	var sysAddressRes model.SysAddressRes
	errResp := json.Unmarshal(body, &sysAddressRes)
	//fmt.Println(json.Unmarshal(body, &sysAddressRes))
	//fmt.Println("--final response-->")
	//fmt.Println(sysAddressRes)

	if errResp != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, sysAddressRes)
}
