package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"

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

func DownloadYoutubeFromUrl(c *gin.Context) {}
