package blockchain

import (
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"github.com/nivjain/golang-gin-web3/api"
	"github.com/nivjain/golang-gin-web3/forms"
)

// UserController will hold everything that controller needs
type BlockchainController struct {
	bapi *api.Api
}

// SetUserController returns a new UserHandler, this stuff is called dependency Injection took a while to understand...
func (h *BlockchainController) SetUserController(conn *api.Api) *BlockchainController {
	return &BlockchainController{
		bapi: conn,
	}
}

//RegisterUser ...
func (h *BlockchainController) GetHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "BlockchainController reporting hi V!"})
}

//GetEmailId ...
func (h *BlockchainController) GetEmailId(c *gin.Context) {
	var userDetails forms.EmailStruct
	var email = "sripal.jain@gmail.com"
	userDetails.EmailID = email
	c.JSON(http.StatusOK, gin.H{"success": &userDetails})
}

//GetMessageSmartContract ...
func (h *BlockchainController) GetMessageSmartContract(c *gin.Context) {
	reply, err := h.bapi.Hello(&bind.CallOpts{})
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//RegisterUser ...
func (h *BlockchainController) GetGreetings(c *gin.Context) {
	message := c.Param("message")
	reply, err := h.bapi.Greet(&bind.CallOpts{}, message)
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}
