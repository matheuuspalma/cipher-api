package main

import (
	"cipher/cmd/server"
	"sync"
)

/*
** Cipher API
** Version 0.0
** schemes http
 */

func main() {

	run := sync.WaitGroup{}
	run.Add(1)

	go func() {
		server.Server()
		run.Done()
	}()

	run.Wait()
}
