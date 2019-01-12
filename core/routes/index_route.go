package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noahp78/CMX/extpoints"
)


func init() {
	fmt.Println("RegisterIndex");
	extpoints.Routes.Register(new (IndexRoute), "index")
}

type IndexRoute struct {}


func (p *IndexRoute) Register(r *gin.Engine){
	fmt.Println("Register Gin ROUTE")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HelloWorld",
		})
	})
}
