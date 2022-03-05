package v1

import (
	"github.com/gin-gonic/gin"
	v1Controllers "github.com/nivjain/golang-gin-web3/controllers/v1/blockchain"
)

//UserRouter ...
type UserRouter struct{}

var uc = new(v1Controllers.UserController)

//HomePage ...
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi Sripal Jain greetings from userroutes homepage",
	})
}

// Routes for user
func (usr UserRouter) Routes(router *gin.Engine) {
	user := router.Group("/v1/user")
	user.GET("/gethome", HomePage)
	user.GET("/gethome2", uc.GetHomePage)
	user.GET("/getemail", uc.GetEmailId)
}
