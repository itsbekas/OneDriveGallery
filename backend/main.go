package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// add your routes here

	r.Run() // listen and serve on 0.0.0.0:8080
}
