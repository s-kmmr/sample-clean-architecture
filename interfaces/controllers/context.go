package controllers

type Context interface {
	Bind(i interface{}) error
	JSON(code int, obj interface{})
}
