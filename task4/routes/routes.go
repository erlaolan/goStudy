package routes

import (
	"github.com/gin-gonic/gin"
	"go_blog/controller"
	"go_blog/middleware"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()

	//使用中间件
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	//创建控制器实例
	authController := controller.AuthController{}
	postController := controller.PostController{}
	commentController := controller.CommentController{}

	//路由
	api := r.Group("/api")
	{
		//认证相关路由（无需认证）
		auth := api.Group("auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			//用户信息
			authenticated.GET("/profile", authController.GetProfile)
			//文章相关路由
			post := authenticated.Group("posts")
			{
				post.POST("", postController.CreatePost)
				post.PUT("/:id", postController.UpdatePost)
				post.DELETE("/:id", postController.DeletePost)
			}
			//评论相关路由
			comment := authenticated.Group("/posts/:post_id/comments")
			{
				comment.POST("", commentController.CreateComment)

			}

		}
		//公开路由（无需认证）
		public := api.Group("")
		{
			public.GET("posts", postController.GetPostList)
			public.GET("posts/:id", postController.GetPost)
		}
		comments := api.Group("comments")
		{
			comments.GET("/post/:post_id", commentController.GetCommentList)

		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Blog is running",
			"status":  "ok",
		})
	})
	return r
}
