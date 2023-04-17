package ef1

import (
	"errors"
	"fmt"
	"sort"

	"rey.com/fairallol/model"
	"rey.com/fairallol/pkg/e"
)

// The code above implements a type of fairness called envy-freeness up to one good (EF1) in its allocation process. An allocation is said to be EF1 if no agent would prefer the bundle of any other agent plus one more item to their own bundle.

// The fairAllocation function first allocates each item to the agent with the highest valuation for that item, which ensures that the allocation is Pareto optimal (no agent can be made better off without making another agent worse off). It then adjusts the allocation to maximize fairness by reallocating an item from the agent with the largest allocation to the agent with the smallest allocation, if there exists an item that the agent with the smallest allocation values more than the agent with the largest allocation. This process aims to make the allocation envy-free, in that no agent would prefer the bundle of any other agent to their own bundle.
// Note that the code does not guarantee the existence of an EF1 allocation for all possible inputs, as the problem of finding such an allocation is known to be NP-hard. However, the code provides a heuristic that attempts to find a good allocation in practice.
// Ensure Pareto optimal but not necessary for min_max
func fairAllocation(agents []*model.Agent, items []*model.Item) (map[string][]string, error) {
	// allocate each item to the agent with the highest valuation for that item
	allocations := make(map[string][]string) // allocations[agent_name] is a slice of item name allocated to this agent

	for _, item := range items {
		fmt.Printf("start allocate %s\n", item.Name)

		highestValuation := -1
		var AgentsWithHighestValuation []*model.Agent

		for _, agent := range agents {
			if _, ok := agent.Allocations[item.Name]; !ok && agent.Valuations[item.Name] > highestValuation {
				highestValuation = agent.Valuations[item.Name]
				AgentsWithHighestValuation = []*model.Agent{agent}
			} else if _, ok := agent.Allocations[item.Name]; !ok && agent.Valuations[item.Name] == highestValuation {
				AgentsWithHighestValuation = append(AgentsWithHighestValuation, agent)
			}
		}

		if len(AgentsWithHighestValuation) > 0 {
			// if there are multiple agents with the same highest valuation, choose the one with the smallest allocation size
			sort.Slice(AgentsWithHighestValuation, func(i, j int) bool {
				return len(allocations[AgentsWithHighestValuation[i].Name]) < len(allocations[AgentsWithHighestValuation[j].Name])
			})
			chosenAgent := AgentsWithHighestValuation[0]
			chosenAgent.Allocations[item.Name] = item
			allocations[chosenAgent.Name] = append(allocations[chosenAgent.Name], item.Name)
			fmt.Printf("allocate %s to %s\n", item.Name, chosenAgent.Name)
		}
	}

	// adjust allocations to maximize fairness
	iteration := 0
	maxDifference := 2
	maxIterations := len(items)
	for iteration = 0; iteration < maxIterations && maxDifference > 1; iteration++ {
		// calculate the allocation size for each agent
		agentSizes := make(map[string]int)
		for _, agent := range agents {
			agentSizes[agent.Name] = len(agent.Allocations)
		}

		// find the agent with the largest allocation size and the agent with the smallest allocation size
		var maxAgent, minAgent *model.Agent
		for _, agent := range agents {
			if maxAgent == nil || agentSizes[agent.Name] > agentSizes[maxAgent.Name] {
				maxAgent = agent
			}
			if minAgent == nil || agentSizes[agent.Name] < agentSizes[minAgent.Name] {
				minAgent = agent
			}
		}

		// calculate the maximum difference between the allocations of any two agents
		maxDifference = agentSizes[maxAgent.Name] - agentSizes[minAgent.Name]

		fmt.Println("max_difference: ", maxDifference)

		// EF1
		if maxDifference > 1 {
			// reallocate an item from the agent with the largest allocation to the agent with the smallest allocation

			// Find an item in the largest agent's allocation that has the highest value to the smallest agent.
			var bestItemInMaxAgent *model.Item
			bestItemInMaxAgentValue := -1

			for _, item := range maxAgent.Allocations {
				minAgentValuation := minAgent.Valuations[item.Name]
				if minAgentValuation >= maxAgent.Valuations[item.Name] && minAgentValuation > bestItemInMaxAgentValue {
					bestItemInMaxAgentValue = minAgent.Valuations[item.Name]
					bestItemInMaxAgent = item
				}
			}

			if bestItemInMaxAgent == nil {
				fmt.Println("No items of A, B prefers")
			} else {
				delete(maxAgent.Allocations, bestItemInMaxAgent.Name)
				minAgent.Allocations[bestItemInMaxAgent.Name] = bestItemInMaxAgent
				allocations[maxAgent.Name] = remove(allocations[maxAgent.Name], bestItemInMaxAgent.Name)
				allocations[minAgent.Name] = append(allocations[minAgent.Name], bestItemInMaxAgent.Name)

				fmt.Printf("swap %s from %s to %s\n", bestItemInMaxAgent.Name, maxAgent.Name, minAgent.Name)

			}

		}
	}

	// format output
	output := make(map[string][]string)

	for agent_name, item_name := range allocations {
		output[agent_name] = item_name
	}

	if iteration == maxIterations {
		return output, errors.New(e.COMPROMISE)
	}

	return output, nil
}

func remove(s []string, elem string) []string {
	for i, e := range s {
		if e == elem {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
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

	allocation, err = fairAllocation(agents, items)

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
