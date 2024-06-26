package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slowhigh/goclean/infra/config"
	"github.com/slowhigh/goclean/infra/router/handler"
	"github.com/slowhigh/goclean/infra/router/middleware"
	"github.com/slowhigh/goclean/internal/controller/rest"
	docs "github.com/slowhigh/goclean/third_party/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	router   *gin.Engine
	restConf config.Rest
}

func NewRouter(conf config.Config, ctrl rest.MemoController) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsHandler())

	docs.SwaggerInfo.Title = conf.Swagger.Title
	docs.SwaggerInfo.Version = conf.Swagger.Version
	docs.SwaggerInfo.Description = conf.Swagger.Description
	docs.SwaggerInfo.Schemes = conf.Swagger.Schemes
	docs.SwaggerInfo.Host = conf.Swagger.Host
	docs.SwaggerInfo.BasePath = conf.Swagger.BasePath

	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		memo := v1.Group("/memo")
		{
			memo.GET("/:id", func(c *gin.Context) { handler.GetMemo(c, ctrl) })
			memo.GET("/", func(c *gin.Context) { handler.ListMemos(c, ctrl) })
			memo.POST("/", func(c *gin.Context) { handler.CreateMemo(c, ctrl) })
			memo.PUT("/:id", func(c *gin.Context) { handler.UpdateMemo(c, ctrl) })
			memo.DELETE("/:id", func(c *gin.Context) { handler.DeleteMemo(c, ctrl) })
		}
	}

	return Router{
		router:   r,
		restConf: conf.Rest,
	}
}

func (r Router) Run() error {
	return r.router.Run(fmt.Sprintf(":%d", r.restConf.Port))
}
