package services

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/host"
	"ota-updater/internal/common"
)

type InfoType struct {
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	KernelArch      string `json:"kernelArch"`
}

func InfoHandler(context *gin.Context) {
	info, err := host.Info()
	if err != nil {
		return
	}
	res := common.Response(0, struct {
		Info InfoType `json:"info"`
	}{
		Info: InfoType{
			Hostname:        info.Hostname,
			OS:              info.OS,
			Platform:        info.Platform,
			PlatformFamily:  info.PlatformFamily,
			PlatformVersion: info.PlatformVersion,
			KernelVersion:   info.KernelVersion,
			KernelArch:      info.KernelArch,
		},
	})
	context.JSON(200, res)
}
