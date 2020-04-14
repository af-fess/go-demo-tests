package core

import (
	coreMocks "github.com/af-fess/go-demo-tests/core/mocks"
	factoriesMocks "github.com/af-fess/go-demo-tests/factories/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

const dummyErrorMessage = "Dummy error"

type DoSuite struct {
	suite.Suite

	testDo *Do

	doerMock *coreMocks.Doer
	systemLibraryMock *factoriesMocks.SystemLibraryProvider
}


func (suite *DoSuite) SetupTest() {
	suite.doerMock = &coreMocks.Doer{}

	suite.systemLibraryMock = &factoriesMocks.SystemLibraryProvider{}

	params := BuildParams{}
	params.IsOptIn = true // default value
	// We do not care about real path.
	// We will check both cases: with error and without
	params.WorkspacePath = "dummy/path"

	suite.testDo = &Do{
		Params: params,
		Doer:     suite.doerMock,
		SystemLibraryProvider: suite.systemLibraryMock,
		os:           suite.systemLibraryMock, // only for tests
		time:         suite.systemLibraryMock, // only for tests
	}
}


func (suite *DoSuite) TearDownTest() {
	suite.doerMock.AssertExpectations(suite.T())
	suite.systemLibraryMock.AssertExpectations(suite.T())
}

func (suite *DoSuite) Test__DoSomething__noError_isCalled() {

	suite.doerMock.On("DoSomething").Return(nil).Once()

	err := suite.testDo.Run()
	suite.NoError(err)
}

//func (suite *DoSuite) Test__DoSomething__Error_isCalled() {
//
//	suite.doerMock.On("DoSomething").Return(errors.NewManage(dummyErrorMessage)).Once()
//
//	err := suite.testDo.Run()
//	suite.EqualError(err, dummyErrorMessage)
//}


func TestDoSuite(t *testing.T) {
	suite.Run(t, new(DoSuite))
}