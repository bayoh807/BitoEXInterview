package v1Router

import (
	"github.com/gin-gonic/gin"
)

type v1Router struct {
}

var V1Router v1Router

const (
	prefix = "/auth"
)

func (r *v1Router) InitRouter(rootG *gin.RouterGroup) {

	r.authGroup(rootG)

}
