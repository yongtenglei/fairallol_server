package randrobin

import (
	"fmt"

	"rey.com/fairallol/model"
)

// agents do not track what they really allocated, output track it
func roundRobin(agents []*model.Agent, items []*model.Item) map[string][]string {
	remainingItems := make([]*model.Item, len(items))
	copy(remainingItems, items)

	allocation := make(map[string][]string)
	for _, agent := range agents {
		allocation[agent.Name] = []string{}
	}

	for {
		for _, agent := range agents {
			if len(remainingItems) == 0 {
				return allocation
			}

			bestItemForAgent := ""
			bestItemValue := 0
			for _, item := range remainingItems {
				itemValue, ok := agent.Valuations[item.Name]
				if !ok {
					continue
				}
				if bestItemForAgent == "" || itemValue > bestItemValue {
					bestItemForAgent = item.Name
					bestItemValue = itemValue
				}
			}

			allocation[agent.Name] = append(allocation[agent.Name], bestItemForAgent)

			for i, item := range remainingItems {
				if item.Name == bestItemForAgent {
					copy(remainingItems[i:], remainingItems[i+1:])
					remainingItems = remainingItems[:len(remainingItems)-1]
					break
				}
			}
		}
	}
}

func RoundRobin(goods []string, agent1Name string, valuation1 map[string]int, agent2Name string, valuation2 map[string]int) (allocation map[string][]string) {
	agents, err := model.InitAgents(goods, agent1Name, valuation1, agent2Name, valuation2)
	if err != nil {
		return nil
	}

	items := model.InitItems(goods)

	fmt.Println("agents: ", agents)
	fmt.Println("items: ", items)

	allocation = roundRobin(agents, items)

	return
}
