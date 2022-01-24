package rest

import (
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.GET("/hello", hello)
	v1.GET("/os", os)
}

func LoadRouter() {
	router := gin.Default()
	routes(router)
	router.Run(":5000")
}
