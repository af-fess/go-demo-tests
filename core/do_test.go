package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DoSuite struct {
	suite.Suite
}


func (suite *DoSuite) SetupTest() {

}


func (suite *DoSuite) TearDownTest() {

}

func TestDoSuite(t *testing.T) {
	suite.Run(t, new(DoSuite))
}