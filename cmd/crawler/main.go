package main

import (
	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/pkg/crawler"
)

func main() {
	cmd.InitBase()

	config.Logger.Info("start crawling danggn")
	crawler.DanggnCrawler()
}
