package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinder-server/providders"
)

func (m *middleware) ResponseMiddleware(c *gin.Context) {

	response := providders.Provider.NewResponse()

	defer func() {
		switch true {
		case response.Data != nil, response.Error != nil:
			code, res := newSchemaResponse(response)
			c.JSON(code, res)
		default:

		}
	}()

	c.Set("response", response)
	c.Next()

}

func newSchemaResponse(resp *providders.Response) (int, map[string]interface{}) {

	httpCode, code, message := func() (int, int, string) {
		if resp.Error == nil {
			return http.StatusOK, 0, "success"
		} else {
			return http.StatusBadRequest, 1, resp.Error.Error()
		}
	}()
	res := map[string]interface{}{
		"code":    code,
		"message": message,
	}

	if code == 0 {
		res["data"] = resp.Data
	}
	if resp.Paging != nil {
		res["paging"] = resp.Paging
	}

	return httpCode, res
}
