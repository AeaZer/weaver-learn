package main

import (
	"github.com/weaver/learn/cmd"
	"log"
)

func main() {
	for _, command := range cmd.RegisteredCommand {
		go func() {
			err := command.Execute()
			if err != nil {
				log.Fatal(err)
				return
			}
		}()
	}
}
