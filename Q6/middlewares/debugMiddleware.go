package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	CurrentENV = func() string {
		return os.Getenv("ENV")
	}()
)

func (m *middleware) DebugMiddleware(c *gin.Context) {

	if mode, has := c.GetQuery("mode"); (CurrentENV == "dev" || CurrentENV == "uat") && has && mode == "debug" {

		var request interface{}
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)
		url := fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.Path)
		method := c.Request.Method
		if json.Valid([]byte(x)) {

			json.Unmarshal([]byte(x), &request)

		} else {

			request = fmt.Sprintf("%s \n", string(x))
		}

		c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
			"url":    url,
			"method": method,
			"query":  c.Request.URL.Query(),
			"body":   request,
		})
		c.Abort()
		return
	} else {
		c.Next()
	}

}
