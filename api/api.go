package api

import (
	"github.com/das-pilot/das-go/das_hyperledger"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/def/fabapi"
	"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	"strconv"
	"encoding/json"
)

var setup = das_hyperledger.WalletSetup{
	ConfigFile:      "conf/wallet-conf.yaml",
	ChannelID:       "daschannel",
	OrgID:           "Org1",
	ChainCodeID:     "wallet",
	ChannelConfig:   "../fabric-1.0-impl/network/channel-artifacts/channel.tx",
	ConnectEventHub: true,
}

var chClient apitxn.ChannelClient

func Init() {
	err := setup.Initialize()
	if err != nil {
		fmt.Printf("Error while initializing %s", err)
	} else {
		fmt.Print("\nApi initialized\n")
	}
	// Create SDK setup for the integration tests
	sdkOptions := fabapi.Options{
		ConfigFile: setup.ConfigFile,
	}

	sdk, err := fabapi.NewSDK(sdkOptions)
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s", err)
	}

	chClient, err = sdk.NewChannelClient(setup.ChannelID, "User1")
	if err != nil {
		fmt.Printf("Failed to create new channel client: %s", err)
	}

	var queryArgs = [][]byte{[]byte("one"), []byte("two")}
	value, err := chClient.Query(apitxn.QueryRequest{ChaincodeID: setup.ChainCodeID, Fcn: "query", Args: queryArgs})
	if err != nil {
		fmt.Printf("Failed to query funds: %s", err)
	} else {
		fmt.Printf("Result: %s", value)
	}
}

//Charge wallet for amount
func Charge(fromWallet, toWallet string, amount float32) CommonResponse {

	if amount <= 0 {
		return CommonResponse{
			Result:  "Failure",
			Message: "Amount to charge must be positive",
		}
	}
	var chargeArgs = [][]byte{[]byte(fromWallet), []byte(toWallet), []byte(strconv.FormatFloat(float64(amount), 'f', 2, 64))}
	_, err := chClient.ExecuteTx(apitxn.ExecuteTxRequest{ChaincodeID: setup.ChainCodeID, Fcn: "charge", Args: chargeArgs})
	if err == nil {
		return CommonResponse{
			Result:  "OK",
			Message: "Charged",
		}
	} else{
		return CommonResponse{
			Result: "Failure",
			Message: fmt.Sprint(err),
		}
	}
}

func ChargeMultiple(fromWallet, toWallet string, amount float32, times int) CommonResponse {
	for i := 0; i < times; i++ {
		resp := Charge(fromWallet, toWallet, amount)
		if resp.Result != "OK" {
			return resp
		}
	}
	return CommonResponse{
		Result:  "OK",
		Message: "Charged",
	}
}

func ChargeMultipleOneBlock(chargesCollection string) CommonResponse {
	var chargeArgs = [][]byte{[]byte(chargesCollection)}
	_, err := chClient.ExecuteTx(apitxn.ExecuteTxRequest{ChaincodeID: setup.ChainCodeID, Fcn: "chargeMultiple", Args: chargeArgs})
	if err == nil {
		return CommonResponse{
			Result:  "OK",
			Message: "Charged",
		}
	} else{
		return CommonResponse{
			Result: "Failure",
			Message: fmt.Sprint(err),
		}
	}
}

func Create(wallet string) CommonResponse {
	var createArgs = [][]byte{[]byte(wallet)}
	_, err := chClient.ExecuteTx(apitxn.ExecuteTxRequest{ChaincodeID: setup.ChainCodeID, Fcn: "create", Args: createArgs})
	if err == nil {
		return CommonResponse{
			Result:  "OK",
			Message: "Created",
		}
	} else{
		return CommonResponse{
			Result: "Failure",
			Message: fmt.Sprint(err),
		}
	}
}

//History called
func History(filter HistoryFilter) []HistoryResponse {
	response := []HistoryResponse{}
	var queryArgs = [][]byte{[]byte(filter.FromWallet), []byte(filter.ToWallet)}
	value, err := chClient.Query(apitxn.QueryRequest{ChaincodeID: setup.ChainCodeID, Fcn: "queryHistory", Args: queryArgs})
	if err != nil {
		response = append(response, HistoryResponse{
			Message: "Failed to query funds",
		})
	} else {
		json.Unmarshal(value, &response)
	}
	return response
}

//Balance called
func Balance(fromWallet, toWallet string) BalanceResponse {
	var queryArgs = [][]byte{[]byte(fromWallet), []byte(toWallet)}
	value, err := chClient.Query(apitxn.QueryRequest{ChaincodeID: setup.ChainCodeID, Fcn: "query", Args: queryArgs})
	if err != nil {
		fmt.Printf("Failed to query funds: %s", err)
		return BalanceResponse{
			Result: "Failure",
			Message: "Failed to query funds",
		}

	} else {
		fmt.Printf("Result: %s", value)
		val, err := strconv.ParseFloat(string(value), 32)
		if err != nil {
			return BalanceResponse{
				Result: "Failure",
				Message: "Failed to parse result",
			}
		} else {
			return BalanceResponse{
				Result:  "OK",
				FromWallet:  fromWallet,
				ToWallet:  toWallet,
				Balance: float32(val),
			}
		}
	}
}
