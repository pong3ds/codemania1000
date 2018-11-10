package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type MemberServiceTestSuite struct {
	suite.Suite
}

func TestMemberServiceTypeSuite(t *testing.T) {
	suite.Run(t, &MemberServiceTestSuite{})
}

func (ts *MemberServiceTestSuite) TestGetMembers_ExpectCorrectMembersReturn() {
	mockRequester := NewMockRequester()
	mockRequester.On("Get", "http://member-service/api/members").Return(`[
		{
			"name": "Chaiyapong"
		},
		{
			"name": "Apaichon"
		}
	]`, nil)

	memberSvc := NewMemberService()
	members, err := memberSvc.getMembers(mockRequester)

	is := assert.New(ts.T())
	if is.NoError(err) {
		is.Equal(2, len(members))
		is.Equal("Chaiyapong", members[0].Name)
		is.Equal("Apaichon", members[1].Name)
	}
	mockRequester.AssertExpectations(ts.T())
}

func (ts *MemberServiceTestSuite) TestGetMembers_GivenError_ExpectErrorReturn() {
	mockRequester := NewMockRequester()
	mockRequester.On("Get", "http://member-service/api/members").
		Return(``, fmt.Errorf("Error 1"))

	memberSvc := NewMemberService()
	members, err := memberSvc.getMembers(mockRequester)

	is := assert.New(ts.T())
	is.Nil(members)
	is.Error(err)
	is.Equal("Error 1", err.Error())

	mockRequester.AssertExpectations(ts.T())
}
