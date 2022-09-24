package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "default.html", gin.H{})
}

func setupRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", defaultHandler)
}

func main() {
	router := gin.Default()
	setupRouter(router)
	err := router.Run(":3000")
	if err != nil {
		log.Fatalf("Run error: %s", err)
	}
}
