package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nivjain/golang-gin-web3/api"
	"github.com/nivjain/golang-gin-web3/blkclient"
	"github.com/nivjain/golang-gin-web3/repositories"

	// mid "github.com/nivjain/golang-gin-web3/middlewares/bcconnection"
	v1 "github.com/nivjain/golang-gin-web3/routers/v1"
)

//Init ...
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

func main() {
	log.Println("heloo")
	port := os.Getenv("PORT")
	EthClientURL := os.Getenv("ETH_CLIENT_URL")
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Recovery())

	blkChainConnection := new(blkclient.Connection)
	conn2 := blkChainConnection.NewConnection()
	// conn3 := blkChainConnection.GetConnection()
	log.Println(conn2.GetConnection())

	client, err := ethclient.Dial(EthClientURL)
	if err != nil {
		panic(err)
	}
	// conn, err := api.NewApi(common.HexToAddress("CONTRACT_ADDRESS"), client)
	// conn, err := api.NewApi(common.HexToAddress("0x9Aa23BcdB51785e79EaFF318E7Db61B0befbd905"), client)
	SmartContractAddress := os.Getenv("SMART_CONTRACT_ADDRESS")
	conn, err := api.NewApi(common.HexToAddress(SmartContractAddress), client)
	if err != nil {
		panic(err)
	}

	// Routes profile
	var userRouter = new(v1.UserRouter)
	var blockchainRouter = new(v1.BlockchainRouter)

	//Init
	// Repo initialization for db and rabbitmq setup and then Routes setup ...
	mCommonClientRepo := repositories.SetCommonClient(conn, client, conn2)
	userRouter.Routes(r)
	blockchainRouter.Routes(mCommonClientRepo, r)
	// blockchainRouter.Routes(conn, client, conn2, r)
	// HTML rendering ...
	r.LoadHTMLGlob("./public/html/*")
	r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})
	// Default route if any of the above routes doesnt match ...
	r.NoRoute(func(c *gin.Context) {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			panic(err)
		}
		c.HTML(404, "404.html", gin.H{"message": "Oops page not found!!!", "dd": reply})
	})
	log.Fatal(r.Run(":" + port))
}

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
