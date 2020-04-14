package main

import (
	"github.com/af-fess/go-demo-tests/core"
	"github.com/af-fess/go-demo-tests/factories"

)

func main() {
	systemLibraryProvide := factories.New()

	params := core.NewBuildParams(
		true,
		"/Users/maximshoustin/AppsFlyer/projects")

	manageInstance := core.NewManage()

	manageInstance.Execute(systemLibraryProvide, params)
}
