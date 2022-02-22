package got

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)


func StartGotServer() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	r := gin.Default()

	r.POST("/router",addRouter)
	r.POST("/sign",addSign)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func addRouter(c *gin.Context) {
	var routerAddress string
	if err := c.BindJSON(&routerAddress); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if routerAddress == "" {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	BroadcastWhiteList.AddRouter(routerAddress)

	c.JSON(http.StatusOK,nil)
}

func addSign(c *gin.Context) {
	var sign string
	if err := c.BindJSON(&sign); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if sign == "" {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	BroadcastWhiteList.AddSign(sign)

	c.JSON(http.StatusOK,nil)
}