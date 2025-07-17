package routers

import (
    "github.com/gin-gonic/gin"
    "lib/internal/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    api := r.Group("/api")
    {
		api.GET("/ping", controllers.Ping)
    }
    return r
}

