package routers

import (
	"final-project/controllers/comment"
	"final-project/controllers/photo"
	"final-project/controllers/socmed"
	"final-project/controllers/user"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", user.RegisterUser)
		userRouter.POST("/login", user.LoginUser)
		userRouter.PUT("/", middleware.Authentication(), user.UpdateUser)
		userRouter.DELETE("/", middleware.Authentication(), user.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", photo.CreatePhoto)
		photoRouter.GET("/", photo.GetAllPhotos)
		photoRouter.PUT("/:photoId", middleware.AuthorizationData("photoId", "photos"), photo.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.AuthorizationData("photoId", "photos"), photo.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", comment.CreateComment)
		commentRouter.GET("/", comment.GetAllComments)
		commentRouter.PUT("/:commentId", middleware.AuthorizationData("commentId", "comments"), comment.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.AuthorizationData("commentId", "comments"), comment.DeleteComment)
	}

	socmedRouter := r.Group("/socialmedias")
	{
		socmedRouter.Use(middleware.Authentication())
		socmedRouter.POST("/", socmed.CreateSocialMedia)
		socmedRouter.GET("/", socmed.GetAllSocialMedia)
		socmedRouter.PUT("/:socialMediaId", middleware.AuthorizationData("socialMediaId", "social_media"), socmed.UpdateSocialMedia)
		socmedRouter.DELETE("/:socialMediaId", middleware.AuthorizationData("socialMediaId", "social_media"), socmed.DeleteSocialMedia)
	}

	return r
}
