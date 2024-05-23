package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistFunc(r *gin.Engine) {

	r.GET("/hi", func(c *gin.Context) {
		c.HTML(http.StatusOK, "../public/page_header.html", gin.H{
			"title": "Admin2 Users",
		})
	})
}
