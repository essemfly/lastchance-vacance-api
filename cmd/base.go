package cmd

import (
	"github.com/1000king/handover/config"
)

func InitBase() {
	config.InitLogger()
	config.InitViper()
	config.InitRegisterRepo()
}
