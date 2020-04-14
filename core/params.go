package core

type BuildParams struct {
	IsOptIn       bool
	WorkspacePath string
}

func NewBuildParams(isOptIn bool, workspacePath string) BuildParams {
	return BuildParams{
		IsOptIn:isOptIn,
		WorkspacePath: workspacePath,
	}
}