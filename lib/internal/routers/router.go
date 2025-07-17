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
		api.GET("/version", controllers.GetYTDLPVersion)
		api.GET("/download", controllers.DownloadYoutubeFromUrl)
    }

    return r
}

