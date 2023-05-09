package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"rey.com/fairallol/allocations/ef1"
	"rey.com/fairallol/model"
	"rey.com/fairallol/pkg/e"
)

// Under Development
func EF1(c *gin.Context) {
	var req model.Reqest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.FAILED,
			"data": err.Error(),
		})
		return
	}

	var res model.Response

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
