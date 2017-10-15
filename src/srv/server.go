package srv

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/das-pilot/das-go/api"
	"github.com/gin-gonic/gin"
)

//Start REST server
func Start(port int) {

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/das/charge", charge)
	router.GET("/das/balance", balance)
	router.GET("/das/history", history)

	router.Run(fmt.Sprintf(":%d", port))
}

func charge(c *gin.Context) {
	wallet := c.PostForm("wallet")
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 32)
	if err == nil {
		result := api.Charge(wallet, float32(amount))
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "Failure", "message": fmt.Sprint(err)})
	}
}

func balance(c *gin.Context) {
	wallet := c.Query("wallet")
	result := api.Balance(wallet)
	c.JSON(http.StatusOK, result)
}

func history(c *gin.Context) {

	wallet := c.Query("wallet")

	timeStart, err := strconv.ParseInt(c.Query("from"), 10, 64)
	if err != nil {
		timeStart = -1
	}

	timeEnd, err := strconv.ParseInt(c.Query("to"), 10, 64)
	if err != nil {
		timeEnd = -1
	}

	filter := api.HistoryFilter{
		Wallet:    wallet,
		TimeStart: timeStart,
		TimeEnd:   timeEnd}

	result := api.History(filter)

	c.JSON(http.StatusOK, result)
}
