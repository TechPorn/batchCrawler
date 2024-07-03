package main

import (
	"batchCrawler/engine"
	"batchCrawler/luofeng/parser"
	"batchCrawler/persist"
	"batchCrawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver("mongodb://localhost:27017", "crawler", "profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //&scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "https://www.lfgvip.com/",
		ParserFunc: parser.ParseProvinceList,
	})
}
