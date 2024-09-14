package main

import (
	item "learnfx/kitex_gen/learn/fx/item/learnfxservice"
	"log"
)

func main() {
	svr := item.NewServer(new(LearnFxServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
