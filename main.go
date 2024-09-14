package main

import (
	"learnfx/domain"
	"learnfx/domain/infra"

	"go.uber.org/fx"
)

// func main() {
// 	svr := item.NewServer(new(LearnFxServiceImpl))

// 	err := svr.Run()

// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// }
func main() {
	fx.New(
		domain.Module,
		infra.Module,
		fx.Provide(NewRpcServer, NewLearnFxServiceImpl),
		fx.Invoke(startServer),
	).Run()
}
