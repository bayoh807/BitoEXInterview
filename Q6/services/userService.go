package services

import (
	"fmt"
	"sync"
	"tinder-server/dto"
	"tinder-server/resource/requests"
	response "tinder-server/resource/responses"
)

type userService struct {
}

var (
	UserService userService
	MatchPool   = map[string]dto.User{}
	UsersPool   = map[string]dto.User{}
)

func (s *userService) GetUser(key string) (*dto.User, error) {

	if user, has := UsersPool[key]; has {
		return &user, nil
	} else {
		return nil, fmt.Errorf("not found user")
	}
}

func (s *userService) CrateUser(req requests.AddUserMatchRequest) *dto.User {
	user := dto.Dto.NewUser(req)
	UsersPool[user.ID] = *user

	return user
}

func (s *userService) AddMatchPool(user *dto.User) error {

	if user.Rule.Times <= 0 {
		return fmt.Errorf("current user's times is zero")
	} else {
		// add to match pool
		MatchPool[user.ID] = *user
		return nil

	}
}

func (s *userService) GetUserFromMatchPool(key string) *dto.User {
	if user, has := MatchPool[key]; !has {
		return nil
	} else {
		return &user
	}
}

/**
* 時間複雜度為 O(n)
**/
func (s *userService) GetMatch(user *dto.User) []interface{} {
	var matchChan chan dto.User = make(chan dto.User)
	matches := make([]interface{}, 0)
	wg := sync.WaitGroup{}
	wg.Add(len(MatchPool))

	for _, item := range MatchPool {
		go func(item dto.User) {
			defer wg.Done() // 在 goroutine 完成时通知 WaitGroup
			if matchUser := s.goMatch(user, &item); matchUser != nil {
				matchChan <- *matchUser
			}
		}(item)
	}

	go func() {
		wg.Wait()
		close(matchChan)
	}()

	for match := range matchChan {
		matches = append(matches, response.UserResponse.NewResource(&match))
	}
	return matches
}

/**
* 時間複雜度為接近 O(1)
**/
func (s *userService) goMatch(currentUser *dto.User, item *dto.User) *dto.User {

	if item.ID != currentUser.ID &&
		// check current user rule
		item.Height >= currentUser.Rule.HeightRange.Start &&
		item.Height <= currentUser.Rule.HeightRange.End &&
		item.Gender == currentUser.Rule.MatchGender &&
		// check match user rule
		currentUser.Height >= item.Rule.HeightRange.Start &&
		currentUser.Height <= item.Rule.HeightRange.End &&
		currentUser.Gender == item.Rule.MatchGender {

		currentUser.Lock.Lock()
		item.Lock.Lock()
		currentUser.Rule.Times -= 1
		item.Rule.Times -= 1
		if item.Rule.Times == 0 {
			s.RemoveMatchPool(item.ID)
		} else {
			item.Lock.Unlock()
			MatchPool[item.ID] = *item
		}

		if currentUser.Rule.Times == 0 {
			s.RemoveMatchPool(currentUser.ID)
		} else {
			currentUser.Lock.Unlock()
			MatchPool[currentUser.ID] = *currentUser
		}

		return item
	}

	return nil

}

func (s *userService) RemoveMatchPool(key string) {
	//TODO : 這邊想確認移除時，如果已被配對到的情境
	delete(MatchPool, key)
}
