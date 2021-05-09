package router

import (
	"github.com/gin-gonic/gin"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/injector"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	i := injector.NewInjector()
	r.GET("/members", func(c *gin.Context) {
		i.MemberController().MemberList(c.Request.Context(), c)
	})
	return r
}
