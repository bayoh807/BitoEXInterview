package v1Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tinder-server/providders"
	"tinder-server/resource/requests"
	response "tinder-server/resource/responses"
	"tinder-server/services"
)

type userController struct {
}

var (
	V1UserController userController
)

func (ctr *userController) AddSinglePersonAndMatch(c *gin.Context) {

	var req requests.AddUserMatchRequest
	resp := c.MustGet("responses").(*providders.Response)

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error = err
		return
	} else {

		user := services.UserService.CrateUser(req)
		if addPoolErr := services.UserService.AddMatchPool(user); addPoolErr != nil {
			resp.Message = addPoolErr.Error()
			return
		} else {
			resp.Data = response.UserResponse.NewResource(user)
		}

	}
}

func (ctr *userController) RemoveSinglePerson(c *gin.Context) {
	var req requests.SingleUserRequest
	resp := c.MustGet("responses").(*providders.Response)

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

	resp := c.MustGet("responses").(*providders.Response)
	userID := c.Param("user_id")

	if _, errGetUser := services.UserService.GetUser(userID); errGetUser != nil {
		resp.Error = fmt.Errorf("not found")
		resp.HttpCode = http.StatusNotFound
		return
	} else if user := services.UserService.GetUserFromMatchPool(userID); user == nil {

		resp.Error = fmt.Errorf("not in the match")
		return
	} else {
		resp.Data = services.UserService.GetMatch(user)
		return
	}
}
