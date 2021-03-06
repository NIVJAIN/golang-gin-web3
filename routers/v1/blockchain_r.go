package v1

import (
	"github.com/gin-gonic/gin"
	v1Controllers "github.com/nivjain/golang-gin-web3/controllers/v1/blockchain"
	"github.com/nivjain/golang-gin-web3/repositories"
)

//BlockchainRouter ...
type BlockchainRouter struct{}

var bcc = new(v1Controllers.BlockchainController)

// Routes for user
func (usr BlockchainRouter) Routes(mCommonClient *repositories.CommonClient, router *gin.Engine) {
	bc := bcc.SetUserController(mCommonClient)
	user := router.Group("/v1/blockchain")
	user.GET("/gethome", HomePage)
	user.GET("/gethome2", bc.GetHomePage)
	user.GET("/getemail", bc.GetEmailId)
	user.GET("/getmessage", bc.GetMessageSmartContract)
	user.GET("/greet/:message", bc.GetGreetings)
	user.GET("/get/counter/count", bc.GetCounterTotalCount)
	user.GET("/counter/increment", bc.CounterIncrement2)
	user.GET("/counter/decrement", bc.CounterDecrement)
}
