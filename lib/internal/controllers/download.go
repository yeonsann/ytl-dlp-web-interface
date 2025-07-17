package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"fmt"

	"os/exec"
)

func GetYTDLPVersion(c *gin.Context) {
	cmdOutput, cmdErr := exec.Command("yt-dlp", "--version").Output()

	if cmdErr != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": string(cmdErr.Error())})
		return 
	}

    c.JSON(http.StatusOK, gin.H{"message": string(cmdOutput)})
}

func DownloadStreamMPD(c *gin.Context) {
	urlParam := c.Query("url") // gets ?url=... from the request

    if urlParam == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "url param is required"})
        return
    }
	
	cmdOutput, cmdErr := exec.Command("yt-dlp", fmt.Sprintf(" '%s'", urlParam)).Output()

	if cmdErr != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": string(cmdErr.Error())})
		return 
	}

    c.JSON(http.StatusOK, gin.H{"message": string(cmdOutput)})
}
