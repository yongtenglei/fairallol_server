package ef1

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFairAllocation(t *testing.T) {
	cases := [][]int{
		{3, 6, 11, 21, 30, 36},
		{2, 6, 11, 21, 31, 35},
		{2, 6, 11, 21, 32, 34},
		{2, 6, 11, 22, 23, 42},
		{9, 12, 16, 18, 25, 26},
		{9, 12, 16, 19, 20, 30},
		{9, 12, 16, 19, 21, 29},
		{9, 12, 16, 19, 22, 28},
		{9, 12, 16, 19, 23, 27},
		{9, 12, 16, 19, 24, 26},
		{9, 12, 16, 20, 21, 28},
		{9, 12, 16, 20, 22, 27},
		{9, 12, 16, 20, 23, 26},
		{9, 12, 16, 20, 24, 25},
		{9, 12, 16, 21, 22, 26},
		{9, 12, 16, 21, 23, 25},
		{9, 12, 16, 22, 23, 24},
		{9, 12, 17, 18, 19, 31},

		{9, 12, 17, 18, 20, 30},
		{3, 10, 12, 14, 23, 44},

		{3, 10, 12, 14, 24, 43},
		{3, 10, 12, 14, 25, 42},
		{3, 10, 12, 14, 26, 41},
		{3, 10, 12, 14, 27, 40},
		{9, 12, 17, 18, 21, 29},
		{9, 12, 17, 18, 22, 28},
		{9, 12, 17, 18, 23, 27},
		{9, 12, 17, 18, 24, 26},
		{9, 12, 17, 19, 20, 29},
		{9, 12, 17, 19, 21, 28},
		{9, 12, 17, 19, 22, 27},
		{9, 12, 17, 19, 23, 26},
		{9, 12, 17, 19, 24, 25},
		{9, 12, 17, 20, 21, 27},
		{9, 12, 17, 20, 22, 26},
		{9, 12, 17, 20, 23, 25},
		{9, 12, 17, 21, 22, 25},
		{9, 12, 17, 21, 23, 24},
		{9, 12, 18, 19, 20, 28},
		{9, 12, 18, 19, 21, 27},
		{9, 12, 18, 19, 22, 26},
		{9, 12, 18, 19, 23, 25},
		{9, 12, 18, 20, 21, 26},
		{9, 12, 18, 20, 22, 25},
		{9, 12, 18, 20, 23, 24},
		{9, 12, 18, 21, 22, 24},
		{9, 12, 19, 20, 21, 25},
		{9, 12, 19, 20, 22, 24},
	}

	items := []*Item{
		{"item1"},
		{"item2"},
		{"item3"},
		{"item4"},
		{"item5"},
		{"item6"},
	}

	counter := 1
	for i, j := 0, 1; i < len(cases); i, j = i+2, j+2 {

		preference1 := make(map[string]int)
		for k := 0; k < len(items); k++ {
			preference1[items[k].name] = cases[i][k]
		}

		preference2 := make(map[string]int)
		for k := 0; k < len(items); k++ {
			preference2[items[k].name] = cases[j][k]
		}

		fmt.Println("=======case " + strconv.Itoa(counter) + " ==========")
		fmt.Println(preference1)
		fmt.Println(preference2)

		agents := []*Agent{
			{"Alice", preference1, make(map[string]*Item)},
			{"Bob", preference2, make(map[string]*Item)},
		}
		fairAllocation(agents, items)
		for _, agent := range agents {
			fmt.Println(agent.name)
			for name := range agent.allocations {
				fmt.Println(name)
			}
		}

		counter++
	}
}
