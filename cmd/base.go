package cmd

import (
	"github.com/1000king/handover/config"
)

func InitBase() {
	config.InitLogger()
	config.InitSlack()
	config.InitViper()
	config.InitRegisterRepo()
	config.InitNotification()
}
