/**
* Welcome to CMX
*
*
* CMX is designed to be the ultimate solution for quickly building next generation websites.
* It's packaged up into multiple core 'addons' that provide the basic functionality
* And community addons that can offer even more.
* CMX also comes with a VUEJS based admin system, and a very basic website to get started with.
*/


//CMX is the complete base package
//go:generate go-extpoints
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/noahp78/CMX/addons"
	"github.com/noahp78/CMX/extpoints"
)

var r *gin.Engine;

func main(){
	r = gin.Default()
	// Register all routes
	for name, route := range extpoints.Routes.All() {
		fmt.Println("Register extension ",name, route);
		route.Register(r);
	}

	r.Run()
}