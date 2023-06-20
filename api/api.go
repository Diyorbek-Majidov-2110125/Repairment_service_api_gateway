package api

import (
	"projects/Repairment_service/Repairment_service_api_gateway/api/handlers"
	"projects/Repairment_service/Repairment_service_api_gateway/config"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	//Order
	r.POST("/createOrder", h.CreateOrder)
	r.GET("/getOrder", h.GetOrderById)
	r.GET("/getAllOrder", h.GetAllOrders)
	r.DELETE("/deleteOrder", h.DeleteOrder)
	r.PATCH("/updateOrder", h.UpdateOrder)

	//User
	r.POST("/createUser", h.CreateUser)
	r.GET("/getUser", h.GetUserById)
	r.GET("/getAllUser", h.GetAllUsers)
	r.DELETE("/deleteUser", h.DeleteUser)
	r.PATCH("/updateUser", h.UpdateUser)

	return
}
