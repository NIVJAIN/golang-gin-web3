package blockchain

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/nivjain/golang-gin-web3/api"
	"github.com/nivjain/golang-gin-web3/forms"
)

// UserController will hold everything that controller needs
type BlockchainController struct {
	bapi      *api.Api
	ethclient *ethclient.Client
}

// SetUserController returns a new UserHandler, this stuff is called dependency Injection took a while to understand...
func (h *BlockchainController) SetUserController(conn *api.Api, ethclient *ethclient.Client) *BlockchainController {
	return &BlockchainController{
		bapi:      conn,
		ethclient: ethclient,
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

//GetGreetings ...
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

//GetCounterTotalCount ...
func (h *BlockchainController) GetCounterTotalCount(c *gin.Context) {
	log.Println("GetCounterTotalCount --- ")
	reply, err := h.bapi.GetCount(&bind.CallOpts{})
	if err != nil {
		// log.Error(err.Error())
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

//GetCounterTotalCount ...
func (h *BlockchainController) CounterIncrement(c *gin.Context) {
	// privateKey, err := crypto.HexToECDSA("")
	private_key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(private_key)

	// privateKey, err := crypto.HexToECDSA("stereo consider quality wild fat farm symptom bundle laundry side one lemon")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := h.ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	// reply, err := h.bapi.Increment(&bind.TransactOpts{})
	reply, err := h.bapi.Increment(auth)
	if err != nil {
		// log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// reply2, err := GetCounterTotalCount()O
	// if err != nil {
	// 	// log.Error(err.Error())
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"success": reply})
}

// //GetCounterTotalCount ...
// func (h *BlockchainController) CounterDecrement(c *gin.Context) {
// 	reply, err := h.bapi.Decrement(&bind.CallOpts{})
// 	if err != nil {
// 		// log.Error(err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": reply})
// }
