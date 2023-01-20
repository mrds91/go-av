package config

import (
	"address-validation/pkg/controller"

	"github.com/gin-gonic/gin"
)

func configureRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/address-validation", controller.AddressValidationController)
	return router
}

func StartApp() {
	router := configureRouter()
	router.Run("mylocal.sprint.com:8080")
}
