package router

import (
	"github.com/gin-gonic/gin"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/injector"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Router struct {
	c *gin.Engine
}

func NewRouter(i injector.Injector) *Router {
	r := gin.Default()

	v1 := r.Group("v1")
	// メンバー一覧の取得
	v1.GET("/members", func(c *gin.Context) {
		i.MemberController().MemberList(c.Request.Context(), c)
	})

	// 新規メンバーの作成
	v1.POST("/members", func(c *gin.Context) {
		i.MemberController().CreateMember(c.Request.Context(), c)
	})

	// swagger setting
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Router{c: r}
}

func (r *Router) Start() {
	if err := r.c.Run(); err != nil {
		panic(err.Error())
	}
}
