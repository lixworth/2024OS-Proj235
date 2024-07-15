package services

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/host"
	"ota-updater/internal/common"
)

type InfoType struct {
	HostName string `json:"hostname" `
}

func InfoHandler(context *gin.Context) {
	info, err := host.Info()
	if err != nil {
		return
	}
	res := common.Response(0, InfoType{HostName: info.Hostname})
	context.JSON(200, res)
}
