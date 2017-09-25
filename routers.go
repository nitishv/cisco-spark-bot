package main

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

func getWebEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}

	return router
}

var routes = []Route{
	Route{"GET", "/", Index},
}
