package model

import "errors"

type Item struct {
	Name string
}

type Agent struct {
	Name        string
	Valuations  map[string]int
	Allocations map[string]*Item
}

func InitAgents(goods []string, agent1Name string, valuation1 map[string]int, agent2Name string, valuation2 map[string]int) ([]*Agent, error) {
	var agents []*Agent

	p1 := make(map[string]int)
	p2 := make(map[string]int)

	for _, good := range goods {
		// p1
		if v1, ok := valuation1[good]; !ok {
			return nil, errors.New("Someone did not make an assessment of all goods")
		} else {
			p1[good] = v1
		}

		// p2
		if v2, ok := valuation2[good]; !ok {
			return nil, errors.New("Someone did not make an assessment of all goods")
		} else {
			p2[good] = v2
		}
	}

	agent1 := &Agent{
		Name:        agent1Name,
		Valuations:  p1,
		Allocations: make(map[string]*Item),
	}

	agent2 := &Agent{
		Name:        agent2Name,
		Valuations:  p2,
		Allocations: make(map[string]*Item),
	}

	agents = append(agents, agent1, agent2)

	return agents, nil
}

func InitItems(goods []string) []*Item {
	var items []*Item

	for _, good := range goods {
		item := &Item{Name: good}
		items = append(items, item)
	}

	return items
}
