package core

import "github.com/af-fess/go-demo-tests/factories"

type Manage struct {
	SystemLibraryProvide *factories.SystemLibraryProvide
}

func (m *Manage) GetSystemLibraryProvider() factories.SystemLibraryProvider {
	return m.SystemLibraryProvide
}

func (m *Manage) Execute() error{

	m.SystemLibraryProvide = &factories.SystemLibraryProvide{}

	params := BuildParams{
		IsOptIn:       true,
		WorkspacePath: "/Users/maximshoustin/AppsFlyer/projects",
	}
	do := Do{
		Params: params,
		SystemLibraryProvider: m.SystemLibraryProvide,
	}
	return do.Run()
}

