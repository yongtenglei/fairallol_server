package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	roundrobin "rey.com/fairallol/allocations/roundRobin"
	"rey.com/fairallol/model"
	"rey.com/fairallol/pkg/e"
)

func RoundRobin(c *gin.Context) {
	var req model.Reqest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.FAILED,
			"data": err.Error(),
		})
		return
	}

	fmt.Println("parameters: ", req.Goods, req.Agent1, req.Valuation1, req.Agent2, req.Valuation2)

	// allocation adjusted winner
	allocation := roundrobin.RoundRobin(req.Goods, req.Agent1, req.Valuation1, req.Agent2, req.Valuation2)
	fmt.Println("before: ", allocation)

	if allocation == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.FAILED,
			"data": nil,
		})
		return
	}

	var res model.Response
	res.Allocation = allocation
	fmt.Println("result: ", res)

	c.JSON(http.StatusOK, gin.H{
		"code": e.OK,
		"data": res,
	})
}
