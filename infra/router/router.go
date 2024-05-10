package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	docs "github.com/slowhigh/goclean/docs"
	"github.com/slowhigh/goclean/infra/config"
	"github.com/slowhigh/goclean/infra/router/handler"
	"github.com/slowhigh/goclean/infra/router/middleware"
	"github.com/slowhigh/goclean/internal/adapter/controller/rest"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	router *gin.Engine
	conf   *config.Config
}

func NewRouter(conf *config.Config, ctrl rest.MemoController) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsHandler())

	docs.SwaggerInfo.BasePath = "/v1/memo"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		memo := v1.Group("/memo")
		{
			memo.GET("/:id", func(c *gin.Context) { handler.GetMemo(c, ctrl) })
			memo.GET("/", func(c *gin.Context) { handler.GetMemos(c, ctrl) })
			memo.POST("/", func(c *gin.Context) { handler.PostMemo(c, ctrl) })
			memo.PUT("/:id", func(c *gin.Context) { handler.PutMemo(c, ctrl) })
			memo.DELETE("/:id", func(c *gin.Context) { handler.DeleteMemo(c, ctrl) })
		}
	}

	return Router{
		router: r,
		conf:   conf,
	}
}

func (r Router) Run() error {
	return r.router.Run(fmt.Sprintf(":%d", r.conf.Rest.Port))
}
