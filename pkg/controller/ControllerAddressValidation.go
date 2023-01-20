package controller

import (
	"address-validation/pkg/model"
	"address-validation/pkg/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddressValidationController(c *gin.Context) {
	var expAddressReq model.ExpAddressReq
	if err := c.BindJSON(&expAddressReq); err != nil {
		var err model.ErrorInfo
		err.Code = "400"
		err.Systemmessage = "parsing error"
		err.Usermessage = "Unknown error while parsing the request"
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if len(expAddressReq.Addresses) < 1 {
		var err model.ErrorInfo
		err.Code = "400"
		err.Systemmessage = "no address found"
		err.Usermessage = "Invalid request - Found no address to validate"
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for i, ai := range expAddressReq.Addresses {
		if strings.TrimSpace(ai.Address.Addressline1) == "" || strings.TrimSpace(ai.Address.ZipCode) == "" || strings.TrimSpace(ai.Address.City) == "" {
			var err model.ErrorInfo
			err.Code = "400"
			err.Systemmessage = "bad request"
			err.Usermessage = "Invalid request - addressline1, postalcode, cityname is mandatory at position - " + strconv.Itoa(i+1)
			c.JSON(http.StatusBadRequest, err)
			return
		}
	}

	service.CallAddressValidationService(c, expAddressReq)
}
