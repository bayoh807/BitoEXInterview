package v1Controller

import (
	"github.com/gin-gonic/gin"
	"tinder-server/providders"
	"tinder-server/resource/requests"
	"tinder-server/services"
)

type userController struct {
}

var (
	V1UserController userController
)

func (ctr *userController) AddSinglePersonAndMatch(c *gin.Context) {

	var req requests.AddUserMatchRequest
	resp := c.MustGet("response").(*providders.Response)

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error = err
		return
	} else {

		user := services.UserService.CrateUser(req)
		if addPoolErr := services.UserService.AddMatchPool(user); addPoolErr != nil {
			resp.Message = addPoolErr.Error()
			return
		}

	}
}

func (ctr *userController) RemoveSinglePerson(c *gin.Context) {
	var req requests.SingleUserRequest
	resp := c.MustGet("response").(*providders.Response)

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error = err
		return
	} else if user, errGetUser := services.UserService.GetUser(req.UserID); errGetUser != nil {
		resp.Error = err
		return
	} else {
		services.UserService.RemoveMatchPool(user.ID)
		return
	}
}

func (ctr *userController) QuerySinglePeople(c *gin.Context) {
	var req requests.SingleUserRequest
	resp := c.MustGet("response").(*providders.Response)

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error = err
		return
	} else if user, errGetUser := services.UserService.GetUser(req.UserID); errGetUser != nil {
		resp.Error = err
		return
	} else {
		services.UserService.RemoveMatchPool(user.ID)
		return
	}
}
