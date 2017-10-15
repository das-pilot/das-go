package api

//Charge wallet for amount
func Charge(wallet string, amount float32) CommonResponse {

	if amount <= 0 {
		return CommonResponse{
			Result:  "Failure",
			Message: "Amount to charge must be positive",
		}
	}

	//TODO: make real charge here
	return CommonResponse{
		Result:  "OK",
		Message: "Charged",
	}
}

//History called
func History(filter HistoryFilter) []HistoryResponse {
	response := []HistoryResponse{}
	for i := 0; i < 10; i++ {
		response = append(response, MockHistoryRow())
	}
	return response
}

//Balance called
func Balance(wallet string) BalanceResponse {

	//Simulate balance retrieval here
	var balance float32 = 100.50

	return BalanceResponse{
		Result:  "OK",
		Wallet:  wallet,
		Balance: balance,
	}

}
