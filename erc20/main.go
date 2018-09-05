package erc20

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/panjiang/go/utils/currencies"
	log "github.com/panjiang/golog"
)

var (
	conn *ethclient.Client
)

func ConnectGeth(host string) {
	var err error
	conn, err = ethclient.Dial(host)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	} else {
		log.Printf("Connected to Geth at: %v\n", host)
	}
}

type BalanceResponse struct {
	Name       string         `json:"name,omitempty"`
	Wallet     common.Address `json:"wallet,omitempty"`
	Symbol     string         `json:"symbol,omitempty"`
	Balance    *big.Int       `json:"balance"`
	EthBalance *big.Int       `json:"eth_balance,omitempty"`
	Decimals   *big.Int       `json:"decimals,omitempty"`
	Block      *types.Block   `json:"block,omitempty"`
	Token      *TokenCaller
}

// TokenContext ERC20 Token实例
type TokenContext struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	Caller   *Token `json:"-"`
}

// Info 信息
func (t *TokenContext) Info() string {
	return fmt.Sprintf("Name: %s, Symbol: %s, Decimals: %d", t.Name, t.Symbol, t.Decimals)
}

// GetBalance 余额
func (t *TokenContext) GetBalance(address string) (*big.Float, error) {
	balanceWei, err := t.Caller.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	balanceStr := BigIntDecimal(balanceWei, t.Decimals)
	log.Debug("GetBalance", t.Symbol, address, balanceStr)
	balanceBig := currencies.FromString(balanceStr)
	return balanceBig, nil
}

// InitToken 初始化合约Token
func InitToken(contract string) (*TokenContext, error) {
	token, err := NewToken(common.HexToAddress(contract), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Token contract: %v\n", err)
		return nil, err
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		log.Printf("Failed to get decimals from contract: %v \n", contract)
		return nil, err
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		log.Printf("Failed to get symbol from contract: %v \n", contract)
		symbol = SymbolFix(contract)
	}

	name, err := token.Name(nil)
	if err != nil {
		log.Printf("Failed to retrieve token name from contract: %v | %v\n", contract, err)
		name = "MISSING"
	}
	return &TokenContext{
		Contract: contract,
		Name:     name,
		Decimals: decimals.Int64(),
		Symbol:   symbol,
		Caller:   token,
	}, nil
}

func GetAccount(contract string, wallet string) (*BalanceResponse, error) {
	var err error
	response := new(BalanceResponse)

	response.Wallet = common.HexToAddress(wallet)

	token, err := NewTokenCaller(common.HexToAddress(contract), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Token contract: %v\n", err)
		return nil, err
	}

	response.Block, err = conn.BlockByNumber(context.TODO(), nil)
	if err != nil {
		log.Printf("Failed to get current block number: %v\n", err)
		response.Block = nil
	}

	response.Decimals, err = token.Decimals(nil)
	if err != nil {
		log.Printf("Failed to get decimals from contract: %v \n", contract)
		return nil, err
	}

	response.EthBalance, err = conn.BalanceAt(context.TODO(), response.Wallet, nil)
	if err != nil {
		log.Printf("Failed to get ethereum balance from address: %v \n", response.Wallet)
	}

	response.Balance, err = token.BalanceOf(nil, response.Wallet)
	if err != nil {
		log.Printf("Failed to get balance from contract: %v %v\n", contract, err)
	}

	response.Symbol, err = token.Symbol(nil)
	if err != nil {
		log.Printf("Failed to get symbol from contract: %v \n", contract)
		response.Symbol = SymbolFix(contract)
	}

	response.Name, err = token.Name(nil)
	if err != nil {
		log.Printf("Failed to retrieve token name from contract: %v | %v\n", contract, err)
		response.Name = "MISSING"
	}

	return response, err
}

func SymbolFix(contract string) string {
	switch common.HexToAddress(contract).String() {
	case "0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0":
		return "EOS"
	}
	return "MISSING"
}

type jsonResponse struct {
	Name       string `json:"name,omitempty"`
	Wallet     string `json:"wallet,omitempty"`
	Symbol     string `json:"symbol,omitempty"`
	Balance    string `json:"balance"`
	EthBalance string `json:"eth_balance,omitempty"`
	Decimals   int64  `json:"decimals,omitempty"`
	Block      int64  `json:"block,omitempty"`
}

func (b *BalanceResponse) Format() *jsonResponse {
	return &jsonResponse{
		b.Name,
		b.Wallet.String(),
		b.Symbol,
		BigIntDecimal(b.Balance, b.Decimals.Int64()),
		BigIntDecimal(b.EthBalance, 18),
		b.Decimals.Int64(),
		b.Block.Number().Int64(),
	}
}

func (b *BalanceResponse) Ok() bool {
	if b.Decimals.Sign() >= 0 && b.Balance.Sign() >= 0 {
		return true
	}
	return false
}

func BigIntDecimal(balance *big.Int, decimals int64) string {
	if balance.Sign() == 0 {
		return "0"
	}
	bal := big.NewFloat(0)
	bal.SetInt(balance)
	pow := BigPow(10, decimals)
	p := big.NewFloat(0)
	p.SetInt(pow)
	bal.Quo(bal, p)
	deci := strconv.Itoa(int(decimals))
	dec := "%." + deci + "f"
	newNum := Clean(fmt.Sprintf(dec, bal))
	return newNum
}

func Clean(newNum string) string {
	stringBytes := bytes.TrimRight([]byte(newNum), "0")
	newNum = string(stringBytes)
	if stringBytes[len(stringBytes)-1] == 46 {
		newNum += "0"
	}
	if stringBytes[0] == 46 {
		newNum = "0" + newNum
	}
	return newNum
}

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
