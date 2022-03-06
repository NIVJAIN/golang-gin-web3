package repositories

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nivjain/golang-gin-web3/forms"
)

//RegisterUser ...
func (h *CommonClient) GetHomePage() string {
	return "CommonClient reporting hi V!"
}

//GetCounterTotalCount ...
func (h *CommonClient) GetCounterTotalCount() (*big.Int, error) {
	log.Println("GetCounterTotalCount --- ")
	reply, err := h.bapi.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reply, nil
}

//GetCounterTotalCount ...
func (h *CommonClient) CounterIncrement2() (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(*h.blkchain.GetPublicKey())
	auth := h.blkchain.GetAuth()
	nonce, err := h.blkchain.GetEclient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := h.ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	reply, err := h.blkchain.GetApi().Increment(auth)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

//GetEmailId ...
func (h *CommonClient) GetEmailId() *forms.EmailStruct {
	var userDetails forms.EmailStruct
	var email = "sripal.jain@gmail.com"
	userDetails.EmailID = email
	return &userDetails
}

//GetMessageSmartContract ...
func (h *CommonClient) GetMessageSmartContract() (string, error) {
	reply, err := h.blkchain.GetApi().Hello(&bind.CallOpts{})
	if err != nil {
		return "nil", err
	}
	return reply, nil
}

//GetGreetings ...
func (h *CommonClient) GetGreetings(message string) (string, error) {
	reply, err := h.blkchain.GetApi().Greet(&bind.CallOpts{}, message)
	if err != nil {
		return "nil", err
	}
	return reply, nil
}

//GetCounterTotalCount ...
func (h *CommonClient) CounterDecrement() (*types.Transaction, error) {
	fromAddress := crypto.PubkeyToAddress(*h.blkchain.GetPublicKey())
	privateKey := h.blkchain.GetPrivateKey()
	chainId := h.blkchain.GetChainId()
	nonce, err := h.blkchain.GetEclient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := h.ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	reply, err := h.blkchain.GetApi().Decrement(auth)
	if err != nil {
		return nil, err
	}
	return reply, nil

}
