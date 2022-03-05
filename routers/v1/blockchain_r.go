package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nivjain/golang-gin-web3/api"
	v1Controllers "github.com/nivjain/golang-gin-web3/controllers/v1/blockchain"
)

//BlockchainRouter ...
type BlockchainRouter struct{}

var bcc = new(v1Controllers.BlockchainController)

// Routes for user
func (usr BlockchainRouter) Routes(conn *api.Api, router *gin.Engine) {
	bc := bcc.SetUserController(conn)
	user := router.Group("/v1/blockchain")
	user.GET("/gethome", HomePage)
	user.GET("/gethome2", bc.GetHomePage)
	user.GET("/getemail", bc.GetEmailId)
	user.GET("/getmessage", bc.GetMessageSmartContract)
	user.GET("/greet/:message", bc.GetGreetings)
}
