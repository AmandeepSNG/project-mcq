package userController

import "github.com/gin-gonic/gin"

func getUserList(ctx *gin.Context) {

}

func getUserDetails(ctx *gin.Context) {

}

func createUser(ctx *gin.Context) {

}

func updateUser(ctx *gin.Context) {

}

func deleteUser(ctx *gin.Context) {}

func RegisterUserRoutes(useRouter *gin.RouterGroup) {
	userRouter := useRouter.Group("/users")

	useRouter.GET("/list", getUserList)
	userRouter.GET("/:userId", getUserDetails)
	userRouter.POST("/", createUser)
	useRouter.PATCH("/:userId", updateUser)
	useRouter.DELETE("/:userId", deleteUser)
}
