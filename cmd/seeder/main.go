package main

import (
	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/pkg/crawler"
	pkg "github.com/1000king/handover/pkg/seeder"
)

func main() {
	cmd.InitBase()
	pkg.AddKeywordSeed(crawler.GlobalStartIndex)
}
