package biz

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L../../../libs -lwallet_core -lgo_mili -Wl,-rpath,libs
#include "WalletCoreMili.h"
*/
import "C"

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
	"tx-builder/internal/conf"

	pb "tx-builder/api/builder/v1"

	"github.com/go-kratos/kratos/v2/log"
	grpc "google.golang.org/grpc"
)

type ChainConfig struct {
	Chain    string   `json:"chain"`
	FullName string   `json:"fullName"`
	Type     string   `json:"type"`
	RpcUrls  []string `json:"rpcURLs"`
	ChainId  int64    `json:"chainId"`
	CoinId   int64    `json:"coinId"`
	NetType  string   `json:"netType"`
	Decimals int64    `json:"decimal"`
}

var ChainConfigMap map[string]ChainConfig
var nodeProxyClient TokenlistClient
var blockSpiderClient TransactionClient
var signerClient WalletClient
var Log *log.Helper

type ChainDataUsecase struct {
	ChainConfigMap    *map[string]ChainConfig
	NodeProxyClient   *TokenlistClient
	BlockSpiderClient *TransactionClient
	SignerClient      *WalletClient
	Log               *log.Helper
}

type RequestEnv struct {
	NodeProxy   string `json:"node_proxy"`
	BlockSpider string `json:"block_spider"`
}

type TokenInfo struct {
	Address        string      `json:"address"`
	Amount         string      `json:"amount"`
	Decimals       int64       `json:"decimals"`
	Symbol         string      `json:"symbol"`
	TokenUri       string      `json:"token_uri,omitempty"`
	CollectionName string      `json:"collection_name"`
	TokenType      string      `json:"token_type,omitempty"`
	TokenId        string      `json:"token_id,omitempty"`
	ItemName       string      `json:"item_name,omitempty"`
	ItemUri        string      `json:"item_uri,omitempty"`
	Tokens         []TokenInfo `json:"tokens,omitempty"`
}

type FeeData struct {
	GasLimit string `json:"gas_limit"`
	GasPrice string `json:"gas_price"`
}

func NewChainData(c *conf.Data, logger log.Logger) *ChainDataUsecase {
	reqEnv, _ := json.Marshal(RequestEnv{
		NodeProxy:   c.NodeProxyServer,
		BlockSpider: c.BlockSpiderServer,
	})
	chainConfig, err := os.ReadFile(c.ChainConfigFile)
	if err != nil {
		panic(err)
	}

	ChainConfigMap = make(map[string]ChainConfig)
	C.chaindata_initChainConfig(C.CString(string(chainConfig)), C.CString("{}"))
	C.GoSetRequestEnv(C.CString(string(reqEnv)))

	var ChainConfigList []ChainConfig
	json.Unmarshal(chainConfig, &ChainConfigList)
	for _, cfg := range ChainConfigList {
		ChainConfigMap[cfg.Chain] = cfg
	}

	proxy, err := grpc.Dial(c.NodeProxyServer, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	nodeProxyClient = NewTokenlistClient(proxy)
	spider, err := grpc.Dial(c.BlockSpiderServer, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	blockSpiderClient = NewTransactionClient(spider)
	signer, err := grpc.Dial(c.WalletServer, grpc.WithInsecure())
	signerClient = NewWalletClient(signer)
	if err != nil {
		panic(err)
	}
	Log = log.NewHelper(logger)
	return &ChainDataUsecase{
		ChainConfigMap:    &ChainConfigMap,
		NodeProxyClient:   &nodeProxyClient,
		BlockSpiderClient: &blockSpiderClient,
		SignerClient:      &signerClient,
	}
}

func GetTransferParams(chain, from, token string, amount big.Int) string {
	// amount 转int
	txParam, _ := json.Marshal(map[string]interface{}{
		"from_address":  from,
		"token_address": token,
		"value":         amount.String(),
	})
	ret := C.GoString(C.chaindata_getTransactionParams(C.CString(chain), C.CString(string(txParam))))
	param := make(map[string]interface{})
	json.Unmarshal([]byte(ret), &param)
	if !param["status"].(bool) {
		return ""
	}
	return param["result"].(string)
}

func BuildTxInput(chain, from, to, token, amount string, txReq *TransactionReq) (string, error) {
	if txReq == nil {
		txReq = &TransactionReq{}
	}

	chainConfig := ChainConfigMap[chain]
	chainType := chainConfig.Type
	amt := ShiftDecimal(amount, int32(chainConfig.Decimals), false)
	tkInfo := &GetTokenInfoResp{}
	if token != "" {
		info, err := nodeProxyClient.GetTokenInfo(context.Background(), &GetTokenInfoReq{
			Data: []*GetTokenInfoReq_Data{&GetTokenInfoReq_Data{
				Chain:   chain,
				Address: token,
			}},
		})
		if err != nil {
			return "", err
		}
		if len(info.Data) != 1 {
			return "", fmt.Errorf("get empty token info: %s:%s", chain, token)
		}
		amt = ShiftDecimal(amount, int32(info.Data[0].Decimals), false)
		tkInfo = info
	}
	params := GetTransferParams(chain, from, token, *amt)
	if len(params) == 0 {
		return "", errors.New("get transaction param error")
	}
	var chainParams map[string]interface{}
	json.Unmarshal([]byte(params), &chainParams)
	tx := map[string]interface{}{}

	if chainType == "TVM" {
		inner := map[string]interface{}{}
		version, _ := strconv.Atoi(chainParams["version"].(string))
		timestamp, _ := strconv.Atoi(chainParams["timestamp"].(string))
		number, _ := strconv.Atoi(chainParams["number"].(string))

		inner["blockHeader"] = map[string]interface{}{
			"txTrieRoot":     Hex2Base64(chainParams["txTrieRoot"].(string)),
			"witnessAddress": Hex2Base64(chainParams["witnessAddress"].(string)),
			"parentHash":     Hex2Base64(chainParams["parentHash"].(string)),
			"version":        version,
			"timestamp":      timestamp,
			"number":         number,
		}
		inner["feeLimit"], _ = strconv.Atoi(chainParams["feeLimit"].(string))
		if token != "" {
			inner["transferTrc20Contract"] = map[string]interface{}{
				"contractAddress": token,
				"ownerAddress":    from,
				"toAddress":       to,
				"amount":          StrNumber2Base64(amount),
			}
		} else {
			_amount, _ := strconv.Atoi(amount)
			inner["transfer"] = map[string]interface{}{
				"ownerAddress": from,
				"toAddress":    to,
				"amount":       _amount,
			}
		}
		tx["transaction"] = inner

	} else if chainType == "EVM" {
		nonce, _ := big.NewInt(0).SetString(chainParams["nonce"].(string), 10)
		gasLimit, _ := big.NewInt(0).SetString(chainParams["gasLimit"].(string), 10)
		gasPrice, _ := big.NewInt(0).SetString(chainParams["gasPrice"].(string), 10)
		tx["chainId"] = Number2Base64(big.NewInt(chainConfig.ChainId))
		tx["gasPrice"] = Number2Base64(gasPrice)
		tx["gasLimit"] = Number2Base64(gasLimit)
		tx["nonce"] = Number2Base64(nonce)
		inner := map[string]interface{}{
			"amount": StrNumber2Base64(amount),
		}
		if token != "" {
			tx["toAddress"] = token
			inner["to"] = to
			tx["transaction"] = map[string]interface{}{
				"erc20_transfer": inner,
			}
		} else {
			tx["toAddress"] = to
			tx["transaction"] = map[string]interface{}{
				"transfer": inner,
			}
		}
		txReq.Nonce = nonce.Int64()
		txReq.FeeAmount = gasLimit.Mul(gasLimit, gasPrice).String()
		feeData, _ := json.Marshal(FeeData{
			GasLimit: gasLimit.String(),
			GasPrice: gasPrice.String(),
		})
		txReq.FeeData = string(feeData)
	}

	txReq.ChainName = chain
	txReq.FromAddress = from
	txReq.ToAddress = to
	txReq.Amount = amt.String()
	if token == "" {
		txReq.TransactionType = "native"
	} else {
		txReq.TransactionType = "transfer"
		txReq.ContractAddress = token
		tokenInfo, _ := json.Marshal(TokenInfo{
			Address:  token,
			Decimals: int64(tkInfo.Data[0].Decimals),
			Amount:   amt.String(),
			Symbol:   tkInfo.Data[0].Symbol,
		})
		txReq.TokenInfo = string(tokenInfo)
	}
	txReq.SendTime = time.Now().Local().Unix()

	txInput, _ := json.Marshal(tx)
	return string(txInput), nil
}

func SendRawTx(chain, rawTx string, txReq *TransactionReq) (string, error) {
	if txReq == nil {
		txReq = &TransactionReq{}
	}
	ret := C.chaindata_sendRawTransaction(C.CString(chain), C.CString(rawTx))
	var txData map[string]interface{}
	json.Unmarshal([]byte(C.GoString(ret)), &txData)
	if !txData["status"].(bool) {
		return "", errors.New(txData["error"].(string))
	}
	txHash := txData["result"].(string)
	txReq.ChainName = chain
	txReq.TransactionHash = txHash
	txReq.Status = "pending"
	resp, err := blockSpiderClient.CreateRecordFromWallet(context.Background(), txReq)
	if err != nil {
		Log.Error(err)
	}
	if !resp.Status {
		Log.Error(resp.Mes)
	}
	return txHash, nil
}

func SendTx(chain, from, to, token, amount, passphrase string) (string, error) {
	var txReq TransactionReq
	txInput, err := BuildTxInput(chain, from, to, token, amount, &txReq)
	if err != nil {
		return "", err
	}
	signResp, err := signerClient.SignTransaction(context.Background(), &SignTransactionRequest{
		Address:    from,
		Passphrase: passphrase,
		TxInput:    txInput,
	})
	if err != nil {
		return "", err
	}
	if signResp.Error != "" {
		return "", errors.New(signResp.Error)
	}
	txReq.SessionId = signResp.SessionId
	txHash, err := SendRawTx(chain, signResp.RawTx, &txReq)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func GetBalance(chain string, assetGroup []*pb.AssetList) ([]*pb.AssetList, error) {
	assetMap := make(map[string]map[string]int32)
	tokenMap := make(map[string]*GetTokenInfoResp_Data)
	for _, assets := range assetGroup {
		for _, asset := range assets.Assets {
			if _, exist := tokenMap[asset.Token]; !exist {
				tokenMap[asset.Token] = nil
			}
		}
	}
	var tokenList []*GetTokenInfoReq_Data
	for token, _ := range tokenMap {
		tokenList = append(tokenList, &GetTokenInfoReq_Data{Chain: chain, Address: token})
	}

	/*get token info*/
	infos, err := nodeProxyClient.GetTokenInfo(context.Background(), &GetTokenInfoReq{Data: tokenList})
	if err != nil {
		return nil, err
	}
	for _, info := range infos.Data {
		tokenMap[info.Address] = info
	}

	/*get balance*/
	for _, assets := range assetGroup {
		if _, exist := assetMap[assets.Owner]; !exist {
			assetMap[assets.Owner] = make(map[string]int32)
		}
		for _, asset := range assets.Assets {
			assetMap[assets.Owner][asset.Token] = int32(tokenMap[asset.Token].Decimals)
		}

	}
	assetMapStr, _ := json.Marshal(assetMap)
	ret := C.GoString(C.chaindata_allBalance(C.CString(chain), C.CString(string(assetMapStr))))
	balanceMap := make(map[string]map[string]string)
	json.Unmarshal([]byte(ret), &balanceMap)
	for _, assets := range assetGroup {
		for _, asset := range assets.Assets {
			asset.Name = tokenMap[asset.Token].Name
			asset.Symbol = tokenMap[asset.Token].Symbol
			if asset.Token == "" || asset.Token == "0x" || asset.Token == assets.Owner {
				asset.Balance = balanceMap[assets.Owner][assets.Owner]
			} else {
				asset.Balance = balanceMap[assets.Owner][asset.Token]
			}
		}
	}
	return assetGroup, nil
}

func Number2Base64(bi *big.Int) string {
	bytes := bi.Bytes()
	if len(bytes) == 0 {
		bytes = []byte{0}
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
func StrNumber2Base64(n string) string {
	bi := new(big.Int)
	bi.SetString(n, 10)
	return Number2Base64(bi)
}
func Hex2Base64(hexString string) string {
	hexString = strings.TrimPrefix(hexString, "0x")
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}
	bytes, _ := hex.DecodeString(hexString)
	return base64.StdEncoding.EncodeToString(bytes)
}

func ShiftDecimal(floatString string, count int32, ceil bool) *big.Int {
	f, _, _ := big.ParseFloat(floatString, 10, 0, big.ToZero)
	pow := new(big.Float).SetInt64(int64(math.Pow(10, float64(count))))
	result := new(big.Float).Mul(f, pow)
	if ceil {
		resultInt, _ := result.Add(result, big.NewFloat(1)).Int(nil)
		return resultInt
	} else {
		resultInt, _ := result.Int(nil)
		return resultInt
	}
}
func Gwei2Wei(floatString string) *big.Int {
	return ShiftDecimal(floatString, 9, false)
}
func ReverseEndian(hexString string) string {
	hexString = strings.TrimPrefix(hexString, "0x")
	if len(hexString)%2 != 0 {
		hexString = "0" + hexString
	}
	bytes, _ := hex.DecodeString(hexString)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return hex.EncodeToString(bytes)
}
