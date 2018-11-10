package main

import "encoding/json"

// MemberService is service for member
type MemberService struct{}

// NewMemberService return new MemberService
func NewMemberService() *MemberService {
	return &MemberService{}
}

func (*MemberService) getMembers(requester IRequester) ([]*Member, error) {
	body, err := requester.Get("http://member-service/api/members")
	if err != nil {
		return nil, err
	}

	members := make([]*Member, 0)
	err = json.Unmarshal([]byte(body), &members)
	if err != nil {
		return nil, err
	}

	return members, nil
}
