package blockchain

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/golang-gin-web3/models"
)

// UserController will hold everything that controller needs
type BlockchainController struct {
	blockchainRepo models.BlockchainRepositoryInterface
}

// SetUserController returns a new UserHandler, this stuff is called dependency Injection took a while to understand...
func (h *BlockchainController) SetUserController(blockchainRepo models.BlockchainRepositoryInterface) *BlockchainController {
	return &BlockchainController{
		blockchainRepo: blockchainRepo,
	}
}

//GetHomePage ...
func (h *BlockchainController) GetHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": h.blockchainRepo.GetHomePage()})
}

//GetCounterTotalCount ...
func (h *BlockchainController) GetCounterTotalCount(c *gin.Context) {
	log.Println("GetCounterTotalCount --- ")
	reply, err := h.blockchainRepo.GetCounterTotalCount()
	if err != nil {
		// log.Error(err.Error())
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//GetCounterTotalCount ...
func (h *BlockchainController) CounterIncrement2(c *gin.Context) {
	log.Println("GetCounterTotalCount --- ")
	reply, err := h.blockchainRepo.CounterIncrement2()
	if err != nil {
		// log.Error(err.Error())
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//GetEmailId ...
func (h *BlockchainController) GetEmailId(c *gin.Context) {
	userDetails := h.blockchainRepo.GetEmailId()
	c.JSON(http.StatusOK, gin.H{"success": &userDetails})
}

//GetMessageSmartContract ...
func (h *BlockchainController) GetMessageSmartContract(c *gin.Context) {
	reply, err := h.blockchainRepo.GetMessageSmartContract()
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//GetGreetings ...
func (h *BlockchainController) GetGreetings(c *gin.Context) {
	message := c.Param("message")
	reply, err := h.blockchainRepo.GetGreetings(message)
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//CounterDecrement ...
func (h *BlockchainController) CounterDecrement(c *gin.Context) {
	reply, err := h.blockchainRepo.CounterDecrement()
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}
