package main

import (
	"log"

	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/pkg/crawler"
)

func main() {
	cmd.InitBase()

	log.Println("start crawling danggn")
	crawler.DanggnCrawler()
}
