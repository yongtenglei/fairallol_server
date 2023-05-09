package model

type Reqest struct {
	Goods      []string       `form:"goods" json:"goods"`
	Agent1     string         `form:"agent1" json:"agent1"`
	Valuation1 map[string]int `form:"valuation1" json:"valuation1"`
	Agent2     string         `form:"agent2" json:"agent2"`
	Valuation2 map[string]int `form:"valuation2" json:"valuation2"`
}
