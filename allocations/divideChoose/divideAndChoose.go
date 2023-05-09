package dividechoose

import (
	"fmt"
	"math"

	"rey.com/fairallol/model"
)

func divideAndChoose(agents []*model.Agent, items []*model.Item) map[string][]string {
	output := make(map[string][]string)

	if len(agents) != 2 {
		return output
	}

	divider, chooser := agents[0], agents[1]
	minDifference := math.MaxInt32
	var bestGroup1, bestGroup2 []string

	// Iterate through all possible item combinations using bitmasks
	for i := 0; i < (1 << len(items)); i++ {
		group1 := []string{}
		group2 := []string{}
		valueGroup1 := 0
		valueGroup2 := 0

		for j, item := range items {
			if i&(1<<j) != 0 {
				group1 = append(group1, item.Name)
				valueGroup1 += divider.Valuations[item.Name]
			} else {
				group2 = append(group2, item.Name)
				valueGroup2 += divider.Valuations[item.Name]
			}
		}

		// update the best Group
		difference := int(math.Abs(float64(valueGroup1 - valueGroup2)))
		if difference < minDifference {
			minDifference = difference
			bestGroup1 = group1
			bestGroup2 = group2
		}
	}

	// chooser chose the best group
	valueBestGroup1 := 0
	valueBestGroup2 := 0
	for _, itemName := range bestGroup1 {
		valueBestGroup1 += chooser.Valuations[itemName]
	}
	for _, itemName := range bestGroup2 {
		valueBestGroup2 += chooser.Valuations[itemName]
	}

	var dividerSet []string
	var chooserSet []string

	if valueBestGroup1 >= valueBestGroup2 {
		chooserSet = bestGroup1
		dividerSet = bestGroup2
		chooser.Allocations = wrap2Map(bestGroup1)
		divider.Allocations = wrap2Map(bestGroup2)
	} else {
		chooserSet = bestGroup2
		dividerSet = bestGroup1
		chooser.Allocations = wrap2Map(bestGroup2)
		divider.Allocations = wrap2Map(bestGroup1)
	}

	output[divider.Name] = dividerSet
	output[chooser.Name] = chooserSet

	return output
}

func wrap2Map(items []string) map[string]*model.Item {
	output := make(map[string]*model.Item)

	for _, item := range items {
		output[item] = &model.Item{Name: item}
	}

	return output
}

func DivideAndChoose(goods []string, agent1Name string, valuation1 map[string]int, agent2Name string, valuation2 map[string]int) (allocation map[string][]string) {
	agents, err := model.InitAgents(goods, agent1Name, valuation1, agent2Name, valuation2)
	if err != nil {
		return nil
	}

	items := model.InitItems(goods)

	fmt.Println("agents: ", agents)
	fmt.Println("items: ", items)

	allocation = divideAndChoose(agents, items)

	return
}
