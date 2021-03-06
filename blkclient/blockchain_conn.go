package blkclient

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nivjain/golang-gin-web3/api"
)

//Connection is the connection created
type Connection struct {
	conn       *api.Api
	eclient    *ethclient.Client
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	auth       *bind.TransactOpts
	chainId    *big.Int
}

//NewConnection...
func (c *Connection) NewConnection() *Connection {
	EthClientURL := os.Getenv("ETH_CLIENT_URL")
	SmartContractAddress := os.Getenv("SMART_CONTRACT_ADDRESS")
	private_key := os.Getenv("PRIVATE_KEY")
	CHAIN_ID := os.Getenv("CHAIN_ID")
	client, err := ethclient.Dial(EthClientURL)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := api.NewApi(common.HexToAddress(SmartContractAddress), client)
	if err != nil {
		panic(err)
	}
	// privateKey, err := crypto.HexToECDSA("stereo consider quality wild fat farm symptom bundle laundry side one lemon")
	privateKey, err := crypto.HexToECDSA(private_key)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	i := new(big.Int)
	b, err := fmt.Sscan(CHAIN_ID, i)
	if err != nil {
		log.Println("error scanning value:", err, b)

	} else {
		fmt.Println(i)
	}

	// chainId := big.NewInt(CHAIN_ID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, i)
	if err != nil {
		log.Fatal(err)
	}
	var d Connection
	d.conn = conn
	d.eclient = client
	d.privateKey = privateKey
	d.publicKey = publicKeyECDSA
	d.auth = auth
	d.chainId = i
	return &d
}

//GetConnection returns the connection which was instantiated
func (c *Connection) GetConnection() *Connection {
	return c
}

func (c *Connection) GetApi() *api.Api {
	return c.conn
}

func (c *Connection) GetEclient() *ethclient.Client {
	return c.eclient
}
func (c *Connection) GetPrivateKey() *ecdsa.PrivateKey {
	return c.privateKey
}
func (c *Connection) GetPublicKey() *ecdsa.PublicKey {
	return c.publicKey
}
func (c *Connection) GetAuth() *bind.TransactOpts {
	return c.auth
}
func (c *Connection) GetChainId() *big.Int {
	return c.chainId
}
