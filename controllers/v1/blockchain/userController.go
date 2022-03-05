package blockchain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nivjain/golang-gin-web3/forms"
)

// UserController will hold everything that controller needs
type UserController struct{}

//RegisterUser ...
func (h *UserController) GetHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "userController reporting hi V!"})
}

//GetEmailId ...
func (h *UserController) GetEmailId(c *gin.Context) {
	var userDetails forms.EmailStruct
	var email = "sripal.jain@gmail.com"
	userDetails.EmailID = email
	c.JSON(http.StatusOK, gin.H{"success": &userDetails})
}
