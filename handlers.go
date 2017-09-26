package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}

func FirstMessages(c *gin.Context) {
	log.Info("GOT POST!")
	log.Info(c.PostForm("data"))
}
