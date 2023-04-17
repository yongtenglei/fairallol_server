package main

import (
	"rey.com/fairallol/router"
)

func main() {
	r := router.SetUp()

	r.Run(":8081")
}
