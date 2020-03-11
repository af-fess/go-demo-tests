package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type DeploySuite struct {
	suite.Suite
}


func (suite *DeploySuite) SetupTest() {

}


func (suite *DeploySuite) TearDownTest() {

}

func TestDeployStageSuite(t *testing.T) {
	suite.Run(t, new(DeploySuite))
}