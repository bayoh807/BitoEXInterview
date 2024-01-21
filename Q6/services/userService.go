package services

import (
	"fmt"
	"tinder-server/dto"
	"tinder-server/resource/requests"
)

type userService struct {
}

var (
	UserService userService
	matchPool   map[string]dto.User
	usersPool   map[string]dto.User
)

func (s *userService) GetUser(key string) (*dto.User, error) {

	if user, has := usersPool[key]; has {
		return &user, nil
	} else {
		return nil, fmt.Errorf("not found user")
	}
}
func (s *userService) CrateUser(req requests.AddUserMatchRequest) *dto.User {
	user := dto.Dto.NewUser(req)
	usersPool[user.ID] = *user
	return user
}

func (s *userService) AddMatchPool(user *dto.User) error {

	if user.Times <= 0 {
		return fmt.Errorf("current user's times is zero")
	} else {
		// add to match pool
		matchPool[user.ID] = *user
		return nil

	}
}

func (s *userService) startMatch(currentUser *dto.User, item *dto.User) (*dto.User, error) {

	if item.ID != currentUser.ID &&
		// check current user rule
		item.Height >= currentUser.Rule.HeightRange.Start &&
		item.Height <= currentUser.Rule.HeightRange.End &&
		item.Gender >= currentUser.Rule.MatchGender &&
		// check match user rule
		currentUser.Height >= item.Rule.HeightRange.Start &&
		currentUser.Height <= item.Rule.HeightRange.End &&
		currentUser.Gender >= item.Rule.MatchGender {

		currentUser.Lock.Lock()
		item.Lock.Lock()
		currentUser.Times -= 1
		item.Times -= 1
		if item.Times <= 0 {
			// item times <= 0 , matchPool this user
			s.RemoveMatchPool(item.ID)
		}
	}

	return nil, fmt.Errorf("not match")

}

func (s *userService) RemoveMatchPool(key string) {
	//TODO : 這邊想確認移除時，如果已被配對到的情境
	delete(matchPool, key)
}
