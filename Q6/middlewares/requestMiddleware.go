package middlewares

import (
	"github.com/gin-gonic/gin"
	"tinder-Server/providders"
)

func (m *middleware) RequestMiddleware(c *gin.Context) {

	request := providders.Provider.NewRequest(c.Request.URL.Query())

	c.Set("request", request)
	c.Next()

}
