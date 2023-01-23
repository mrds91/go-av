package controller

import (
	"address-validation/pkg/service"
	"address-validation/pkg/util"
	"fmt"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache"
	"github.com/gin-gonic/gin"
)

// cache init
var CartCache = ttlcache.NewCache()

// Enable global ttl
const ItemExpireWithGlobalTTL time.Duration = 0

func UpdateCartController(c *gin.Context) {

	CartId := c.Param("cartId")
	fmt.Println("cartid-->set-->" + CartId)
	AddressValidationController(c)

	if service.SystemError != nil {
		c.JSON(500, service.SystemError)
	} else {
		jsonData := service.SysAddressSuccessRes
		fmt.Println(jsonData)
		CartCache.SetWithTTL(CartId, jsonData, time.Duration(time.Duration(util.ContextProperty.Cachettl)*time.Second))
	}

}

func FetchCartDetails(c *gin.Context) {

	CartId := c.Param("cartId")
	value, exists := CartCache.Get(CartId)
	fmt.Println(exists)
	if exists == true {
		fmt.Printf("cache found value: %v\n", value)
		c.JSON(http.StatusOK, value)

		return
	} else {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "No details available for CartId:" + CartId,
		})
	}

}
