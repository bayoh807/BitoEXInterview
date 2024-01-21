package dto

import (
	"github.com/google/uuid"
	"sync"
	"tinder-server/resource/requests"
)

type User struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Height uint64   `json:"height"`
	Gender uint32   `json:"gender"`
	Rule   UserRule `json:"rule"`
	Lock   sync.Mutex
}

type UserRule struct {
	HeightRange struct {
		Start uint64
		End   uint64
	}
	MatchGender uint32
	Times       uint64 `json:"times"`
}

func (d *dto) NewUser(req requests.AddUserMatchRequest) *User {

	return &User{
		ID:     uuid.NewString(),
		Name:   req.Name,
		Height: req.Height,
		Gender: *req.Gender,
		Rule: UserRule{
			HeightRange: struct {
				Start uint64
				End   uint64
			}{Start: req.Rule.Range.Start, End: req.Rule.Range.End},
			MatchGender: *req.Rule.Gender,
			Times:       req.Rule.Times,
		},
	}
}
