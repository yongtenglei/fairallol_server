package adjustedwinner

import (
	"fmt"
	"math"
	"math/rand"

	"rey.com/fairallol/model"
)

const threshold = 1000

// Based on Brams and King's approach.
// The item with the highest point value is assigned to the person who ranked it highest,
// and then that person is removed from consideration.
// Repeat two steps above until all items have been assigned.
// *********************************************************
// To ensure fairness, Brams and King's approach uses an adjustment factor to ensure that each person is assigned items of roughly equal value.
// The adjustment factor = sum of points / number of agents.
// Use an annealing mechanism to ensure that allocation proceeds smoothly. (Not getting stuck in a dead-end loop)
// adjusted value = original valuation - (adjustment factor * Number of items owned)
// *********************************************************
// PERF: The current implementation uses a brute force approach to check if each agent has been assigned a set of items they value at least as highly as any other agent's assigned items. A more efficient approach would be to use a linear programming solver to optimize the allocation.
func aw(agents []*model.Agent, items []*model.Item) map[string][]string {
	// Calculate the sum of the point values of all the items
	totalPoints := 0
	for _, agent := range agents {
		for _, valuation := range agent.Valuations {
			totalPoints += valuation
		}
	}

	// The adjustment factor = sum of points / number of agents.
	adjustmentFactor := float64(totalPoints) / float64(len(agents))
	if adjustmentFactor <= 0 {
		panic("adjustmentFactor <= 0")
	}

	allocatedItems := make(map[string]bool)

	// Shuffle the order of the items
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	counter := 1
	for {
		if counter > threshold {
			rate := float64(threshold) / float64(counter)
			adjustmentFactor *= rate
		}

		// Shuffle the order of the agents
		rand.Shuffle(len(agents), func(i, j int) {
			agents[i], agents[j] = agents[j], agents[i]
		})

		allocateItemsToAgents(agents, items, allocatedItems, adjustmentFactor)

		done := checkAndReallocateItems(agents, items, allocatedItems)

		counter++
		if done {
			break
		}
	}

	// format output
	output := make(map[string][]string)

	for _, agent := range agents {
		goods := make([]string, 0)
		for name := range agent.Allocations {
			goods = append(goods, name)
		}
		output[agent.Name] = goods
	}

	return output
}

func allocateItemsToAgents(agents []*model.Agent, items []*model.Item, allocatedItems map[string]bool, adjustmentFactor float64) {
	for _, item := range items {
		// Find the agent with the highest adjusted valuation for this item
		var bestAgent *model.Agent
		bestAdjustedValuation := math.Inf(-1)

		// if already allocated, so do not propose this item again.
		for _, agent := range agents {
			if allocatedItems[item.Name] {
				continue
			}

			// Calculate the agent's adjusted valuation for this item
			adjustedValuation := float64(agent.Valuations[item.Name]) - adjustmentFactor*float64(len(agent.Allocations))
			// Check if this agent has the highest adjusted valuation so far
			if adjustedValuation > bestAdjustedValuation {
				bestAgent = agent
				bestAdjustedValuation = adjustedValuation
			}
		}

		// Assign the item to the best agent (if any)
		if bestAgent != nil {
			bestAgent.Allocations[item.Name] = item
			allocatedItems[item.Name] = true
			// fmt.Println("================")
			// fmt.Println("item: ", item)
			// fmt.Println("bestAgent: ", bestAgent.Name)
			// fmt.Println("================")
		} else {
			// fmt.Println("================")
			// fmt.Println("item: ", item)
			// fmt.Println("cannot find a bestAgent this term")
			// fmt.Println("================")
		}
	}
}

// Check if each agent has been assigned a set of items they value at least as highly as any other agent's assigned items.
func checkAndReallocateItems(agents []*model.Agent, items []*model.Item, allocatedItems map[string]bool) bool {
	done := true
	for _, agent1 := range agents {
		for _, agent2 := range agents {
			if agent1 == agent2 {
				continue
			}

			sum1 := 0
			sum2 := 0
			// fmt.Printf("agent1: %s, agent2: %s\n", agent1.Name, agent2.Name)
			// fmt.Printf("agent1: %d, agent2: %d\n", sum1, sum2)

			for _, item := range items {
				allocation1 := agent1.Allocations[item.Name]
				allocation2 := agent2.Allocations[item.Name]
				valuation1 := agent1.Valuations[item.Name]
				valuation2 := agent2.Valuations[item.Name]
				// fmt.Println("item: ", item)
				// fmt.Println("allocation1: ", allocation1)
				// fmt.Println("valuation1: ", valuation1)
				// fmt.Println("allocation2: ", allocation2)
				// fmt.Println("valuation2: ", valuation2)
				if allocation1 != nil {
					// fmt.Printf("sum1 %d + %d = ", sum1, valuation1)
					sum1 += valuation1
					// fmt.Printf("%d\n", sum1)
					if allocation1 != allocation2 {
						// fmt.Printf("sum2 %d + %d = ", sum2, valuation2)
						sum2 += valuation2
						// fmt.Printf("%d\n", sum2)
					}
				}
			}

			// fmt.Println("agents:")
			// fmt.Printf("agent1: %#v\n", agent1.Allocations)
			// fmt.Printf("agent2: %#v\n", agent2.Allocations)

			// fmt.Println("sum1:sum2", sum1, sum2)

			if sum1 < sum2 {
				done = false

				// Reallocate items assigned to agent1
				reallocateItems(agent1, items, allocatedItems)

				// fmt.Println("after clean:")
				// fmt.Printf("agent1: %#v\n", agent1.Allocations)
				// fmt.Printf("agent2: %#v\n", agent2.Allocations)
			}
		}
	}

	return done
}

func reallocateItems(agent *model.Agent, items []*model.Item, allocatedItems map[string]bool) {
	for _, item := range items {
		if agent.Allocations[item.Name] != nil {
			delete(agent.Allocations, item.Name)
			allocatedItems[item.Name] = false
		}
	}
}

func AdjustedWinner(goods []string, agent1Name string, valuation1 map[string]int, agent2Name string, valuation2 map[string]int) (allocation map[string][]string) {
	agents, err := model.InitAgents(goods, agent1Name, valuation1, agent2Name, valuation2)
	if err != nil {
		return nil
	}

	items := model.InitItems(goods)

	fmt.Println("agents: ", agents)
	fmt.Println("items: ", items)

	allocation = aw(agents, items)

	return
}
