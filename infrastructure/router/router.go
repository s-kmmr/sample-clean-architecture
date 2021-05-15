package router

import (
	"github.com/gin-gonic/gin"
	"github.com/s-kmmr/sample-clean-architecture/infrastructure/injector"
)

type Router struct {
	c *gin.Engine
}

func NewRouter(i injector.Injector) *Router {
	r := gin.Default()
	// i := injector.NewInjector()
	r.GET("/members", func(c *gin.Context) {
		i.MemberController().MemberList(c.Request.Context(), c)
	})
	return &Router{c: r}
}

func (r *Router) Start() {
	if err := r.c.Run(); err != nil {
		panic(err.Error())
	}
}
