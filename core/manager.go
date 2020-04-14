package core

import (
	"github.com/af-fess/go-demo-tests/factories"
)

type Manage struct {}

func NewManage() *Manage {
	return &Manage{}
}


func (m *Manage) Execute(systemLibraryProvide *factories.SystemLibraryProvide, params BuildParams) error{

	do := Do{
		Params: params,
		SystemLibraryProvider: systemLibraryProvide,
	}

	do.Doer = &do

	return do.Run()
}

