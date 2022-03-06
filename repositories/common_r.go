package repositories

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nivjain/golang-gin-web3/api"
	"github.com/nivjain/golang-gin-web3/blkclient"
)

// CommonClient implements models.UserRepository
type CommonClient struct {
	bapi      *api.Api
	ethclient *ethclient.Client
	blkchain  *blkclient.Connection
}

// SetComminClient ..
func SetCommonClient(conn *api.Api, ethclient *ethclient.Client, blkchain *blkclient.Connection) *CommonClient {
	return &CommonClient{
		bapi:      conn,
		ethclient: ethclient,
		blkchain:  blkchain,
	}
}

// GetCommonClient ...
func (m *CommonClient) GetCommonClient() *CommonClient {
	return m
}
