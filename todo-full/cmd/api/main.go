package main

import "github.com/gin-gonic/gin"

func main() {

	var router *gin.Engine = gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "todo api is running",
			"status":  "success",
		})
	})

	router.Run(":3000")
}
