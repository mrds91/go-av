package config

import (
	"address-validation/pkg/controller"
	"address-validation/pkg/util"

	"github.com/gin-gonic/gin"
)

func configureRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/address-validation", controller.AddressValidationController)
	router.POST("/cart/:cartId", controller.UpdateCartController)
	router.GET("/cart/:cartId", controller.FetchCartDetails)
	return router
}

func StartApp() {
	router := configureRouter()
	util.LoadProperty()
	//ContextProperty := properties.MustLoadFile("./pkg/config/app.properties", properties.UTF8)
	router.Run(util.ContextProperty.AppHost + ":" + util.ContextProperty.AppPort)
}
