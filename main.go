package main

import (
	"sync"

	"github.com/ascenttree/jj-go/common"
	"github.com/ascenttree/jj-go/crossarea"
	"github.com/ascenttree/jj-go/update"
)

func main() {
	updateServer := update.NewUpdateServer(
		"0.0.0.0",
		8000,
		common.CreateLogger("update", common.DEBUG),
	)

	crossareaServer := crossarea.NewCrossareaServer(
		"0.0.0.0",
		9511,
		common.CreateLogger("crossarea", common.DEBUG),
	)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		updateServer.Serve()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		crossareaServer.Serve()
	}()

	wg.Wait()
}
