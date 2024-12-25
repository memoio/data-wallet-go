package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/data-wallet-go/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	com "github.com/memoio/contractsv2/common"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
)

type Controller struct {
	endpoint      string
	privateKey    *ecdsa.PrivateKey
	didTransactor *bind.TransactOpts
	proxyAddr     common.Address
	logger        *log.Helper
}

func NewMfileDIDController(chain string, logger *log.Helper) (*Controller, error) {
	instanceAddr, endpoint := com.GetInsEndPointByChain(chain)

	client, err := ethclient.DialContext(context.TODO(), endpoint)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.Error(err)
		chainID = big.NewInt(985)
	}

	privateKey, err := crypto.HexToECDSA(config.Privatekey)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// new instanceIns
	instanceIns, err := inst.NewInstance(instanceAddr, client)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	proxyAddr, err := instanceIns.Instances(&bind.CallOpts{}, com.TypeDidProxy)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	auth.Value = big.NewInt(0) // in wei

	return &Controller{
		endpoint:      endpoint,
		privateKey:    privateKey,
		didTransactor: auth,
		proxyAddr:     proxyAddr,
		logger:        logger,
	}, nil
}
