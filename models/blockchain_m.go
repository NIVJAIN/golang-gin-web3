package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nivjain/golang-gin-web3/forms"
)

// UserRepository ...
type BlockchainRepositoryInterface interface {
	GetHomePage() string
	GetCounterTotalCount() (*big.Int, error)
	CounterIncrement2() (*types.Transaction, error)
	GetEmailId() *forms.EmailStruct
	GetMessageSmartContract() (string, error)
	GetGreetings(message string) (string, error)
	CounterDecrement() (*types.Transaction, error)
}
