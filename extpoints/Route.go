package extpoints

import "github.com/gin-gonic/gin"

type Route interface {
	Register(r *gin.Engine)
}
