package response

import "github.com/s-kmmr/sample-clean-architecture/domain/model"

type MemberResponse struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

func NewMemberResponses(m []model.Member) []MemberResponse {
	if len(m) < 1 {
		return nil
	}
	res := make([]MemberResponse, len(m))
	for i, v := range m {
		res[i] = MemberResponse{
			LastName:  v.LastName(),
			FirstName: v.FirstName(),
		}
	}
	return res
}
