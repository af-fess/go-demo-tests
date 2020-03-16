package core

import (
	"errors"
	"fmt"
	"github.com/af-fess/go-demo-tests/factories"
	"github.com/af-fess/go-demo-tests/utils"
	"os/exec"
	"time"
)

type Doer interface {
	DoSomething() error
}

type Do struct {
	Doer
	Params BuildParams
	SystemLibraryProvider factories.SystemLibraryProvider
	os factories.SystemLibraryProvider
	time factories.SystemLibraryProvider
}

func (d *Do) buildPrettySystemApi() {
	d.os = d.SystemLibraryProvider
	d.time = d.SystemLibraryProvider
}

func (d *Do) Run() error {

	d.buildPrettySystemApi()

	if d.Params.IsOptIn == true {
		err := d.Doer.DoSomething()
		if err != nil {
			return err
		}
	} else {
		return errors.New("DoSomethingWithException")
	}
	return nil
}

func (d *Do) DoSomething() error {

	err := d.os.Chdir(d.Params.WorkspacePath)
	if err != nil {
		fmt.Println("Cannot change a dir into", d.Params.WorkspacePath)
		return err
	}

	c := exec.Command("pwd")

	output, err := utils.RunCommandWithOutput(c)
	if err != nil {
		return err
	}
	fmt.Println("pwd output: ", output)

	d.time.Sleep(5 * time.Second)

	c = exec.Command("ls")
	output, err = utils.RunCommandWithOutput(c)
	if err != nil {
		return err
	}
	fmt.Println("ls -l output: ", output)

	return nil
}
