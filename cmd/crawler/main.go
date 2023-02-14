package main

import (
	"time"

	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/pkg/crawler"
)

func main() {
	cmd.InitBase()
	for {
		crawler.DanggnCrawler()
		time.Sleep(30 * time.Minute)
	}
}
