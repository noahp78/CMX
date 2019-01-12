package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/noahp78/CMX/extpoints"
)


func init() {
	extpoints.Routes.Register(new (IndexRoute), "index")
}

type IndexRoute struct {}


func (p *IndexRoute) Register(r *gin.Engine){
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HelloWorld",
		})
	})
}
