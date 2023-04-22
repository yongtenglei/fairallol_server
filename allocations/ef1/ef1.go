package ef1

import (
	"fmt"

	"rey.com/fairallol/model"
)

// The code above implements a type of fairness called envy-freeness up to one good (EF1) in its allocation process. An allocation is said to be EF1 if no agent would prefer the bundle of any other agent plus one more item to their own bundle.

// The ef1 function first allocates each item to the agent with the highest valuation for that item, which ensures that the allocation is Pareto optimal (no agent can be made better off without making another agent worse off). It then adjusts the allocation to maximize fairness by reallocating an item from the agent with the largest allocation to the agent with the smallest allocation, if there exists an item that the agent with the smallest allocation values more than the agent with the largest allocation. This process aims to make the allocation envy-free, in that no agent would prefer the bundle of any other agent to their own bundle.
// Note that the code does not guarantee the existence of an EF1 allocation for all possible inputs, as the problem of finding such an allocation is known to be NP-hard. However, the code provides a heuristic that attempts to find a good allocation in practice.
// Ensure Pareto optimal but not necessary for min_max
func ef1(agents []*model.Agent, items []*model.Item) map[string][]string {
	allocations := make(map[string][]string) // allocations[agent_name] is a slice of item name allocated to this agent
	unallocatedItems := make([]*model.Item, len(items))
	copy(unallocatedItems, items)

	// Loop until no unallocated items left
	for len(unallocatedItems) > 0 {
		envyFreeAgent := make(map[string]bool)
		for _, agent := range agents {
			envyFreeAgent[agent.Name] = true
			for _, otherAgent := range agents {
				if agent != otherAgent && !isEnvyFreeUpToOneItem(agent, otherAgent, allocations) {
					envyFreeAgent[agent.Name] = false
					break
				}
			}
		}

		for _, item := range unallocatedItems {
			highestValuation := -1
			var AgentsWithHighestValuation []*model.Agent

			for _, agent := range agents {
				// Check if agent has already been envy-free for one item
				if envyFreeAgent[agent.Name] {
					continue
				}

				if agent.Valuations[item.Name] > highestValuation {
					highestValuation = agent.Valuations[item.Name]
					AgentsWithHighestValuation = []*model.Agent{agent}
				} else if agent.Valuations[item.Name] == highestValuation {
					AgentsWithHighestValuation = append(AgentsWithHighestValuation, agent)
				}
			}

			if len(AgentsWithHighestValuation) > 0 {
				chosenAgent := AgentsWithHighestValuation[0]
				allocations[chosenAgent.Name] = append(allocations[chosenAgent.Name], item.Name)
				envyFreeAgent[chosenAgent.Name] = true
			}
		}

		// Remove allocated items from unallocated items list
		for agentName, isEnvyFree := range envyFreeAgent {
			if isEnvyFree {
				for _, item := range allocations[agentName] {
					unallocatedItems = removeItem(unallocatedItems, item)
				}
			}
		}
	}

	// format output
	output := make(map[string][]string)

	for agentName, itemName := range allocations {
		output[agentName] = itemName
	}

	return output
}

func removeItem(items []*model.Item, itemName string) []*model.Item {
	for i, item := range items {
		if item.Name == itemName {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}

func isEnvyFreeUpToOneItem(agent *model.Agent, otherAgent *model.Agent, allocations map[string][]string) bool {
	totalAgentValuation := 0
	totalOtherAgentValuation := 0
	highestValuation := -1

	for _, itemName := range allocations[otherAgent.Name] {
		totalAgentValuation += agent.Valuations[itemName]
		totalOtherAgentValuation += otherAgent.Valuations[itemName]

		if otherAgent.Valuations[itemName] > highestValuation {
			highestValuation = otherAgent.Valuations[itemName]
		}
	}

	// Check if agent i envies agent j's allocation after removing the most valuable item
	return totalAgentValuation >= totalOtherAgentValuation-highestValuation
}

// you should double check both values returned
// if allocation == nil -> internal error -> check error for more detail
// else err != nil -> check if it is `COMPROMISE`, if so, it is fine.
// else bad thing...
func EF1(goods []string, agent1Name string, valuation1 map[string]int, agent2Name string, valuation2 map[string]int) (allocation map[string][]string, err error) {
	agents, err := model.InitAgents(goods, agent1Name, valuation1, agent2Name, valuation2)
	if err != nil {
		return nil, err
	}

	items := model.InitItems(goods)

	fmt.Println("agents: ", agents)
	fmt.Println("items: ", items)

	allocation = ef1(agents, items)

	return
}

//func main() {
//items := []*Item{
//{"item1"},
//{"item2"},
//{"item3"},
//{"item4"},
//{"item5"},
//{"item6"},
//{"item7"},
//}
//items = []*Item{
//{"item1"},
//{"item2"},
//{"item3"},
//{"item4"},
//}
//agents := []*Agent{
//{"Alice", map[string]int{"item1": 6, "item2": 5, "item3": 4, "item4": 3, "item5": 2, "item6": 1}, make(map[string]*Item)},
//{"Bob", map[string]int{"item1": 1, "item2": 2, "item3": 3, "item4": 4, "item5": 5, "item6": 6}, make(map[string]*Item)},
//}

//agents = []*Agent{
//{"Alice", map[string]int{"item1": 1, "item2": 2, "item3": 3, "item4": 4, "item5": 5, "item6": 6}, make(map[string]*Item)},
//{"Bob", map[string]int{"item1": 1, "item2": 2, "item3": 3, "item4": 4, "item5": 5, "item6": 6}, make(map[string]*Item)},
//}

////agents = []*Agent{
////{"Alice", map[string]int{"item1": 1, "item2": 4, "item3": 3, "item4": 4, "item5": 5, "item6": 6}, make(map[string]*Item)},
////{"Bob", map[string]int{"item1": 1, "item2": 2, "item3": 3, "item4": 6, "item5": 5, "item6": 6}, make(map[string]*Item)},
////}

////agents = []*Agent{
////{"Alice", map[string]int{"item1": 9, "item2": 12, "item3": 17, "item4": 18, "item5": 20, "item6": 30}, make(map[string]*Item)},
////{"Bob", map[string]int{"item1": 3, "item2": 10, "item3": 12, "item4": 19, "item5": 23, "item6": 44}, make(map[string]*Item)},
////}
//// *****************itme7
////agents = []*Agent{
////{"Alice", map[string]int{"item1": 9, "item2": 12, "item3": 17, "item4": 18, "item5": 20, "item6": 30, "item7": 13}, make(map[string]*Item)},
////{"Bob", map[string]int{"item1": 3, "item2": 10, "item3": 12, "item4": 19, "item5": 23, "item6": 44, "item7": 8}, make(map[string]*Item)},
////}

////agents = []*Agent{
////{"Alice", map[string]int{"item1": 9, "item2": 12, "item3": 17, "item4": 18, "item5": 20, "item6": 30, "item7": 8}, make(map[string]*Item)},
////{"Bob", map[string]int{"item1": 3, "item2": 10, "item3": 12, "item4": 14, "item5": 23, "item6": 44, "item7": 8}, make(map[string]*Item)},
////}

//agents = []*Agent{
//{"Alice", map[string]int{"item1": 1, "item2": 2, "item3": 3, "item4": 4}, make(map[string]*Item)},
//{"Bob", map[string]int{"item1": 4, "item2": 3, "item3": 2, "item4": 1}, make(map[string]*Item)},
//}

//agents = []*Agent{
//{"Alice", map[string]int{"item1": 1, "item2": 4, "item3": 3, "item4": 2}, make(map[string]*Item)},
//{"Bob", map[string]int{"item1": 4, "item2": 2, "item3": 3, "item4": 1}, make(map[string]*Item)},
//}

//m, err := FairAllocation(agents, items)
//if err != nil {
//fmt.Println(err.Error())
//}

//fmt.Printf("#%v\n", m)
//// fmt.Printf("#%v\n", FairAllocation(agents, items))
//}
