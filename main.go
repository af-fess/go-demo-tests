package main

import (
	"fmt"
	"github.com/af-fess/go-demo-tests/core"
)

func main() {

	manage := core.Manage{}
    err := manage.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
