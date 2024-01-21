package v1Router

import (
	"github.com/gin-gonic/gin"
	"tinder-server/controllers/v1Controller"
)

const (
	authPrefix = "/users"
)

func (r *v1Router) authGroup(rg *gin.RouterGroup) {
	pg := rg.Group(authPrefix)
	{
		pg.GET("single/:user_id", v1Controller.V1UserController.QuerySinglePeople)
		pg.POST("addAndMatch", v1Controller.V1UserController.AddSinglePersonAndMatch)
		pg.DELETE("match", v1Controller.V1UserController.RemoveSinglePerson)
	}

}
