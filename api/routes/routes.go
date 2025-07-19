package routes

import (
	"first_server/api/controllers"
	"first_server/helper"

	"first_server/validations"

	"first_server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	commonRoutes := r.Group("/")
	{
		commonRoutes.GET("/", func(ctx *gin.Context) {
			helper.ResponseHandler(ctx, http.StatusOK, "Server is online", nil)
		})
	}

	// all versioning APIs

	v1 := r.Group("/v1")
	{
		usersRoute := v1.Group("/users")
		{
			usersRoute.GET("/", controllers.ListUsers)
			usersRoute.GET("/:id", controllers.GetUser)
			usersRoute.POST("/", controllers.CreateUser)
			usersRoute.PATCH("/", controllers.UpdateUser)
			usersRoute.DELETE("/", controllers.DeleteUser)
		}

		authRoute := v1.Group("/auth")
		{
			authRoute.POST("/login", middleware.Validator[validations.LoginStruct](), controllers.Login)
			authRoute.POST("/register", middleware.Validator[validations.RegisterStruct](), controllers.Register)
		}
	}
}
