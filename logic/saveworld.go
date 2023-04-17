package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rey.com/fairallol/allocations/ef1"
	"rey.com/fairallol/pkg/e"
)

type SaveWorldReq struct {
	Goods      []string       `form:"goods" json:"goods"`
	Agent1     string         `form:"agent1" json:"agent1"`
	Valuation1 map[string]int `form:"valuation1" json:"valuation1"`
	Agent2     string         `form:"agent2" json:"agent2"`
	Valuation2 map[string]int `form:"valuation2" json:"valuation2"`
}

type SaveWorldRes struct {
	Allocation map[string][]string
}

func SaveWorld(c *gin.Context) {
	var req SaveWorldReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.FAILED,
			"data": err.Error(),
		})
		return
	}

	var res SaveWorldRes

	fmt.Println("parameters: ", req.Goods, req.Agent1, req.Valuation1, req.Agent2, req.Valuation2)

	// allocation ef1
	allocation, err := ef1.EF1(req.Goods, req.Agent1, req.Valuation1, req.Agent2, req.Valuation2)
	fmt.Println("before: ", allocation)
	if err != nil {
		// A COMPROMISE is ok
		if err.Error() == e.COMPROMISE {
			res.Allocation = allocation
			fmt.Println("result: ", res)

			c.JSON(http.StatusOK, gin.H{
				"code": e.OK,
				"data": res,
			})
			return

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": e.FAILED,
				"data": err.Error(),
			})
			return
		}
	}

	res.Allocation = allocation
	fmt.Println("result: ", res)

	c.JSON(http.StatusOK, gin.H{
		"code": e.OK,
		"data": res,
	})
}
