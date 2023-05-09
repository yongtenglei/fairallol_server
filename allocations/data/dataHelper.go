package data

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"rey.com/fairallol/model"
)

type DataFile struct {
	N         int
	Pattern   string
	SampleNum int
	FileName  string
	FileData  [][]int
}

func ReadDataFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersStr := strings.Split(line, ",")
		numbers := make([]int, len(numbersStr))
		for i, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			numbers[i] = num
		}
		data = append(data, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func CalculateScore(agents []*model.Agent, allocation map[string][]string) map[string]int {
	score := make(map[string]int)

	// Initialize scores to 0
	for _, agent := range agents {
		score[agent.Name] = 0
	}

	// Calculate scores
	for agentName, itemNames := range allocation {
		agent := findAgentByName(agentName, agents)
		for _, itemName := range itemNames {
			score[agent.Name] += agent.Valuations[itemName]
		}
	}

	return score
}

// Helper function to find an agent by name in a slice of agents
func findAgentByName(name string, agents []*model.Agent) *model.Agent {
	for _, agent := range agents {
		if agent.Name == name {
			return agent
		}
	}
	return nil
}

func CalulateGoodsDiff(allocation map[string][]string) int {
	values := []int{}

	for _, goods := range allocation {
		values = append(values, len(goods))
	}

	return int(math.Abs(float64(values[0] - values[1])))
}
