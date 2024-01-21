package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tinder-server/dto"
	"tinder-server/resource/requests"
	"tinder-server/services"
)

func TestGetUser_Success(t *testing.T) {

	expectedUser := dto.User{
		ID: "123",
	}

	s := services.UserService
	services.UsersPool[expectedUser.ID] = expectedUser
	// Execution
	result, err := s.GetUser("123")

	// Validation
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, *result)
}

func TestGetUser_NotFound(t *testing.T) {

	expectedUser := dto.User{
		ID: "456",
	}

	s := services.UserService
	services.UsersPool[expectedUser.ID] = expectedUser

	// Execution
	result, err := s.GetUser("666")

	// Validation
	assert.Nil(t, result)
	assert.Equal(t, err.Error(), "not found user")
}

func TestCreateUser_Success(t *testing.T) {

	gender := uint32(0)
	rule := requests.MatchRule{
		Gender: &gender,
		Times:  10,
	}
	rule.Range.Start = 160
	rule.Range.End = 180
	req := &requests.AddUserMatchRequest{
		Name:   "A",
		Height: 180,
		Gender: &gender,
		Rule:   rule,
	}

	user := services.UserService.CrateUser(*req)

	// Execution
	result, has := services.UsersPool[user.ID]

	// Validation

	assert.True(t, has)
	assert.Equal(t, user.ID, result.ID)
}

func TestAddMatchPool_Success(t *testing.T) {

	rule := dto.UserRule{
		Times: 10,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	services.UserService.AddMatchPool(expectedUser)

	// Execution
	result, has := services.MatchPool[expectedUser.ID]

	// Validation
	assert.True(t, has)
	assert.Equal(t, expectedUser.ID, result.ID)
}

func TestAddMatchPool_Fail(t *testing.T) {

	rule := dto.UserRule{
		Times: 0,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	// Execution
	err := services.UserService.AddMatchPool(expectedUser)

	// Validation
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "current user's times is zero")
}

func TestGetUserFromMatchPool_Success(t *testing.T) {

	rule := dto.UserRule{
		Times: 10,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	services.MatchPool[expectedUser.ID] = *expectedUser

	// Execution
	user := services.UserService.GetUserFromMatchPool(expectedUser.ID)

	// Validation
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
}

func TestGetUserFromMatchPool_Fail(t *testing.T) {

	rule := dto.UserRule{
		Times: 10,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	services.MatchPool[expectedUser.ID] = *expectedUser

	// Execution
	user := services.UserService.GetUserFromMatchPool("456")

	// Validation
	assert.Nil(t, user)
}

func TestRemoveMatchPool_Success(t *testing.T) {

	rule := dto.UserRule{
		Times: 10,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	services.MatchPool[expectedUser.ID] = *expectedUser

	// Execution
	services.UserService.RemoveMatchPool(expectedUser.ID)

	_, has := services.MatchPool[expectedUser.ID]
	// Validation
	assert.False(t, has)
}

func TestRemoveMatchPool_Fail(t *testing.T) {

	rule := dto.UserRule{
		Times: 10,
	}
	expectedUser := &dto.User{
		ID:   "123",
		Rule: rule,
	}

	services.MatchPool[expectedUser.ID] = *expectedUser

	// Execution
	services.UserService.RemoveMatchPool("456")

	user, has := services.MatchPool[expectedUser.ID]
	// Validation
	assert.True(t, has)
	assert.Equal(t, expectedUser.ID, user.ID)
}

func TestGetMatch_GetMatch(t *testing.T) {

	ruleA := dto.UserRule{
		MatchGender: 0,
		Times:       10,
		HeightRange: struct {
			Start uint64
			End   uint64
		}{Start: 160, End: 180},
	}
	expectedUserA := &dto.User{
		ID:     "456",
		Gender: 1,
		Height: 170,
		Name:   "A",
		Rule:   ruleA,
	}
	services.MatchPool[expectedUserA.ID] = *expectedUserA

	ruleB := dto.UserRule{
		MatchGender: 1,
		Times:       10,
		HeightRange: struct {
			Start uint64
			End   uint64
		}{Start: 160, End: 180},
	}
	expectedUserB := &dto.User{
		ID:     "123",
		Gender: 0,
		Height: 160,
		Name:   "B",
		Rule:   ruleB,
	}

	services.MatchPool[expectedUserB.ID] = *expectedUserB

	// Execution
	result := services.UserService.GetMatch(expectedUserA)

	// Validation
	assert.Equal(t, 1, len(result))
	assert.Equal(t, expectedUserB.ID, result[0].(map[string]interface{})["id"])
}

func TestGetMatch_NotGetMatch(t *testing.T) {

	ruleA := dto.UserRule{
		MatchGender: 1,
		Times:       10,
		HeightRange: struct {
			Start uint64
			End   uint64
		}{Start: 160, End: 180},
	}
	expectedUserA := &dto.User{
		ID:     "456",
		Gender: 1,
		Height: 170,
		Name:   "A",
		Rule:   ruleA,
	}
	services.MatchPool[expectedUserA.ID] = *expectedUserA

	ruleB := dto.UserRule{
		MatchGender: 1,
		Times:       10,
		HeightRange: struct {
			Start uint64
			End   uint64
		}{Start: 160, End: 180},
	}
	expectedUserB := &dto.User{
		ID:     "123",
		Gender: 0,
		Height: 160,
		Name:   "B",
		Rule:   ruleB,
	}

	services.MatchPool[expectedUserB.ID] = *expectedUserB

	// Execution
	result := services.UserService.GetMatch(expectedUserA)

	// Validation
	assert.Equal(t, 0, len(result))
}
