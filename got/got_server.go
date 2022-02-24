package got

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGotServer() {
	r := gin.Default()

	r.POST("/router", addRouter)
	r.POST("/sign", addSign)

	r.Run() // listen and serve on 0.0.0.0:8080
}

type RouterParam struct {
	Router string `json:"router"`
}
type SignParma struct {
	Sign string `json:"sign"`
}

func addRouter(c *gin.Context) {
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

func addSign(c *gin.Context) {
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
