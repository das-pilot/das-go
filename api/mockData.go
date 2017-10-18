package api

import (
	"math/rand"
	"time"
)

//MockHistoryRow - generates random history row
func MockHistoryRow() HistoryResponse {
	wallet := randStringRunes(16)
	amount := rand.Float32()*200 - 100
	message := "INCOME"
	if amount < 0 {
		message = "SPENT"
	}
	timeStamp := randomTimestamp().Format("15.01.2006 15:04:05")

	return HistoryResponse{
		Wallet:  wallet,
		Amount:  amount,
		Message: message,
		Time:    timeStamp,
	}
}

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterRunes = []rune("0123456789ABCDEF")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomTimestamp() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow
}
