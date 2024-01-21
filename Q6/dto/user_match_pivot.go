package dto

import (
	"sync"
	"tinder-server/resource/requests"
)

type MatchPool struct {
	UserID string
	Gender uint32
	Range  struct {
		Start uint64
		End   uint64
	}
	MatchGender uint32
	Times       uint64
	Lock        sync.Mutex
}

func (d *dto) NewMatchPool(user *User, rule requests.MatchRule) *MatchPool {

	return &MatchPool{
		UserID: user.ID,
		Range: struct {
			Start uint64
			End   uint64
		}{Start: rule.Range.Start, End: rule.Range.End},
		Gender:      user.Gender,
		MatchGender: rule.Gender,
		Times:       rule.Times,
	}
}
