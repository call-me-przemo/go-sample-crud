package main

import (
	"github.com/call-me-przemo/go-sample-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run()
}
