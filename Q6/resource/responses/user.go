package response

import "tinder-server/dto"

type userResponse struct {
}

var UserResponse userResponse

func (r *userResponse) NewResource(user *dto.User) map[string]interface{} {

	return map[string]interface{}{
		"id":     user.ID,
		"name":   user.Name,
		"gender": user.Gender,
		"height": user.Height,
		"rule":   user.Rule,
	}
}
