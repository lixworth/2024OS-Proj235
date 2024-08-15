package services

import (
	"ota-updater/internal/common"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(context *gin.Context) {
	common.ExecCommand("/bin/bash write_image_by_dd.sh")
}

func CheckVersionHandler(context *gin.Context) {
	common.ExecCommand("/bin/bash sysupdate.sh")
}

func ManualUpdateHandler(context *gin.Context) {
	common.ExecCommand("/bin/bash write_image_by_dd.sh")
}

func CheckStatusHandler(context *gin.Context) {
	common.ExecCommand("/bin/bash osfullinfo.sh")
}
