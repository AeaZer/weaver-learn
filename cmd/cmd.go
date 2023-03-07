package cmd

import (
	"github.com/weaver/learn/cmd/components"
	"github.com/weaver/learn/cmd/simple"
)

type Command interface {
	Execute() error
}

var RegisteredCommand = []Command{
	&simple.Command{},
	&components.Command{},
}
