package main

import (
	"fmt"
	"github.com/af-fess/go-demo-tests/core"
)

func main() {

	params := core.BuildParams{
		IsOptIn:       true,
		WorkspacePath: "/Users/maximshoustin/AppsFlyer/projects",
	}
	deploy := core.Deploy{Params: params}
	err := deploy.Run()
	if err != nil {
		fmt.Println(err)
	}
}
