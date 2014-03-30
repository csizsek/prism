package main

import "github.com/csizsek/prism/config"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/dispatcher"
import "github.com/csizsek/prism/frontend"
import "github.com/csizsek/prism/backend"

func main() {
	config := config.ParseConfig()
	front2disp := make(chan *entity.CommonEntity)
	disp2back := make(chan *entity.CommonEntity)
	frontend := frontend.NewFrontend(config.FrontendConfig, front2disp)
	backend := backend.NewBackend(config.Backends[0], disp2back)
	endpoints := []chan *entity.CommonEntity{disp2back}
	dispatcher := dispatcher.NewDispatcher(front2disp, endpoints)

	(*backend).Start()
	go dispatcher.Dispatch()
	(*frontend).Start()

	quit := make(chan int)
	_ = <-quit
}
