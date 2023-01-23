package service

import (
	"address-validation/pkg/model"
	"address-validation/pkg/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var SysAddressSuccessRes model.SysAddressRes
var SystemError error

func CallAddressValidationService(c *gin.Context, expAddressReq model.ExpAddressReq) {

	var sysAddressInfo model.SysAddressInfo
	sysAddressInfo.Addressline1 = expAddressReq.Addresses[0].Address.Addressline1
	sysAddressInfo.Addressline2 = expAddressReq.Addresses[0].Address.Addressline2
	sysAddressInfo.Cityname = expAddressReq.Addresses[0].Address.City
	sysAddressInfo.Statecode = expAddressReq.Addresses[0].Address.State
	sysAddressInfo.Postalcode = expAddressReq.Addresses[0].Address.ZipCode
	pbuff := new(bytes.Buffer)
	json.NewEncoder(pbuff).Encode(sysAddressInfo)

	av_url := "http://" + util.ContextProperty.BackendHost + "/" + util.ContextProperty.BackendPath
	req, err := http.NewRequest("POST", av_url, pbuff)

	if err != nil {
		fmt.Println(err.Error())
	}
	//Headers
	req.Header.Add("x-rapidapi-key", "YOU_API_KEY")
	req.Header.Add("Content-Type", "application/json")
	//timeout settings
	var myClient = &http.Client{Timeout: 10 * time.Second}
	res, SysErr := myClient.Do(req)

	if SysErr != nil {
		SystemError = SysErr
		fmt.Print(SysErr.Error())
		return
	}

	//parse the response
	defer res.Body.Close()
	body, readerr := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if readerr != nil {
		fmt.Println(readerr.Error())
		fmt.Println(body)
	}

	var sysAddressRes model.SysAddressRes

	errResp := json.Unmarshal(body, &sysAddressRes)

	if errResp != nil {
		return
	}
	SysAddressSuccessRes = sysAddressRes
	c.IndentedJSON(http.StatusCreated, sysAddressRes)
}
