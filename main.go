package main

import (
	"github.com/blockfint/di-example-go/app/cmd"
	"github.com/blockfint/di-example-go/app/config"
)

func init() {
	config.Initialize()
}

func main() {
	cmd.Execute()
}
