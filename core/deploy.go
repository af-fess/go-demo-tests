package core

import (
	"errors"
	"fmt"
	"github.com/af-fess/go-demo-tests/utils"
	"os"
	"os/exec"
	"time"
)

type Deploy struct {
	Params BuildParams
}

func (s *Deploy) Run() error {
	if s.Params.IsOptIn == true {
		err := s.DoSomething()
		if err != nil {
			return err
		}
	} else {
		return errors.New("DoSomethingWithException")
	}
	return nil
}

func (s *Deploy) DoSomething() error {

	err := os.Chdir(s.Params.WorkspacePath)
	if err != nil {
		fmt.Println("Cannot change a dir into", s.Params.WorkspacePath)
		return err
	}

	c := exec.Command("pwd")

	output, err := utils.RunCommandWithOutput(c)
	if err != nil {
		return err
	}
	fmt.Println("pwd output: ", output)

	time.Sleep(5 * time.Second)

	c = exec.Command("ls")
	output, err = utils.RunCommandWithOutput(c)
	if err != nil {
		return err
	}
	fmt.Println("ls -l output: ", output)

	return nil
}
