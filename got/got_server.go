package got

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGotServer() {
	r := gin.Default()

	r.POST("/broadcast/router", broadcastAddRouter)
	r.POST("/broadcast/sign", broadcastAddSign)

	r.POST("/filter/to", filterAddTo)
	r.POST("/filter/sign", filterAddSign)

	r.GET("/filter/to", filterListTo)
	r.GET("/filter/sign", filterListSign)

	r.GET("/filter/enable", enable)
	r.GET("/filter/disable", disable)

	r.Run() // listen and serve on 0.0.0.0:8080
}

type RouterParam struct {
	Router string `json:"router"`
}
type ToParam struct {
	To string `json:"to"`
}
type SignParma struct {
	Sign string `json:"sign"`
}

func broadcastAddRouter(c *gin.Context) {
	var routerParam = RouterParam{}
	if err := c.BindJSON(&routerParam); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if routerParam.Router == "" {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	BroadcastWhiteList.AddRouter(routerParam.Router)

	c.JSON(http.StatusOK, "success")
}

func broadcastAddSign(c *gin.Context) {
	var signParam = SignParma{}
	if err := c.BindJSON(&signParam); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if len(signParam.Sign) < 4 {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	BroadcastWhiteList.AddSign(signParam.Sign)

	c.JSON(http.StatusOK, "success")
}

func filterAddTo(c *gin.Context) {
	var toParam = ToParam{}
	if err := c.BindJSON(&toParam); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if toParam.To == "" {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	PendingTxFilter.AddTo(toParam.To)

	c.JSON(http.StatusOK, "success")
}

func filterAddSign(c *gin.Context) {
	var signParam = SignParma{}
	if err := c.BindJSON(&signParam); err != nil {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
		return
	}

	if len(signParam.Sign) < 4 {
		c.JSON(http.StatusBadRequest, "params unmarshal failed")
	}

	PendingTxFilter.AddSign(signParam.Sign)

	c.JSON(http.StatusOK, "success")
}

func filterListTo(c *gin.Context) {

	data, err := PendingTxFilter.ListTo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, string(data))
}

func filterListSign(c *gin.Context) {

	data, err := PendingTxFilter.ListSign()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, string(data))
}

func enable(c *gin.Context) {
	PendingTxFilter.SetEnable()
	c.JSON(http.StatusOK, string("success"))
}

func disable(c *gin.Context) {
	PendingTxFilter.SetDisable()
	c.JSON(http.StatusOK, string("success"))
}
