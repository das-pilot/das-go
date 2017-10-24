package api

//HistoryFilter - structure to filter history records
type HistoryFilter struct {
	FromWallet string
	ToWallet   string
	TimeStart  int64
	TimeEnd    int64
}

//BalanceResponse - balance response holder
type BalanceResponse struct {
	Result  string
	Message string
	FromWallet  string
	ToWallet  string
	Balance float32
}

//CommonResponse - basic response structure
type CommonResponse struct {
	Result  string
	Message string
}

//HistoryResponse - basic response for single history row
type HistoryResponse struct {
	Wallet  string
	Amount  float32
	Message string
	Time    string
}
