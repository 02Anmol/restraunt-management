package routes

import (
	controller "golang-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetorderItem())
	incomingRoutes.POST("/orderItems", controller.CreateorderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateorderItems())
}
