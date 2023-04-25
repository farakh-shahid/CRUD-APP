package routes

import (
	"github.com/farakh-shahid/CRUD-APP/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, uc *controllers.UserController) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PUT("/update/:id", uc.UpdateUser)
	userroute.DELETE("/delete/:id", uc.DeleteUser)
}
