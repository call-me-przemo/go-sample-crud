package routes

import (
	"github.com/call-me-przemo/go-sample-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", controllers.Index)
}
