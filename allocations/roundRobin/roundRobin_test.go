package randrobin

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"rey.com/fairallol/allocations/data"
	"rey.com/fairallol/model"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// comprehensive test
func TestRoundRobin(t *testing.T) {
	fmt.Println("===============Test roundRobin=================")
	dataPath := "../data"

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// prepare for data
		if filepath.Ext(file.Name()) == ".txt" {
			parts := strings.Split(file.Name(), "_")
			if len(parts) >= 4 {
				n, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatal(err)
				}
				pattern := parts[2]
				sampleNum, err := strconv.Atoi(strings.TrimSuffix(parts[3], ".txt"))
				if err != nil {
					log.Fatal(err)
				}

				fileData, err := data.ReadDataFile(filepath.Join(dataPath, file.Name()))
				if err != nil {
					log.Fatal(err)
				}

				dataFile := data.DataFile{
					N:         n,
					Pattern:   pattern,
					SampleNum: sampleNum,
					FileName:  file.Name(),
					FileData:  fileData,
				}

				testRoundRobin(dataFile)
			}
		}
	}
}

func testRoundRobin(fileData data.DataFile) {
	tracker := data.Tracker{
		ScoresDiff: make([]int, 0),
		NumCases:   len(fileData.FileData),
	}

	fmt.Printf("Test for:\n\tN: %d, Pattern: %s, SampleNum: %d, FileName: %s, FileDataLen: %d\n", fileData.N, fileData.Pattern, fileData.SampleNum, fileData.FileName, len(fileData.FileData))

	cases := fileData.FileData
	N := fileData.N

	items := make([]*model.Item, N)
	for i := 0; i < N; i++ {
		items[i] = &model.Item{Name: "item" + strconv.Itoa(i+1)}
	}

	// counter := 1
	tracker.StartTime = time.Now()

	// test process
	if fileData.Pattern == data.TIE {
		for i := 0; i < len(cases); i++ {

			preference1 := make(map[string]int)
			preference2 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference1[items[k].Name] = cases[i][k]
				preference2[items[k].Name] = cases[i][k]
			}

			// fmt.Println("=======case " + strconv.Itoa(counter) + " ==========")
			// fmt.Println(preference1)
			// fmt.Println(preference2)

			agents := []*model.Agent{
				{"Alice", preference1, make(map[string]*model.Item)},
				{"Bob", preference2, make(map[string]*model.Item)},
			}

			allocation := roundRobin(agents, items)

			// fmt.Println(allocation)
			score := data.CalculateScore(agents, allocation)

			// fmt.Println("Scores:")
			values := []int{}
			for _, s := range score {
				// fmt.Printf("%s: %d\n", agentName, s)
				values = append(values, s)
			}

			tracker.ScoresDiff = append(tracker.ScoresDiff, int(math.Abs(float64(values[0]-values[1]))))
			tracker.GoodsDiff = append(tracker.GoodsDiff, data.CalulateGoodsDiff(allocation))

			// counter++

		}
	} else {
		for i, j := 0, 1; i < len(cases); i, j = i+2, j+2 {

			preference1 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference1[items[k].Name] = cases[i][k]
			}

			preference2 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference2[items[k].Name] = cases[j][k]
			}

			// fmt.Println("=======case " + strconv.Itoa(counter) + " ==========")
			// fmt.Println(preference1)
			// fmt.Println(preference2)

			agents := []*model.Agent{
				{"Alice", preference1, make(map[string]*model.Item)},
				{"Bob", preference2, make(map[string]*model.Item)},
			}

			allocation := roundRobin(agents, items)

			// fmt.Println(allocation)
			score := data.CalculateScore(agents, allocation)

			// fmt.Println("Scores:")
			values := []int{}
			for _, s := range score {
				// fmt.Printf("%s: %d\n", agentName, s)
				values = append(values, s)
			}

			tracker.ScoresDiff = append(tracker.ScoresDiff, int(math.Abs(float64(values[0]-values[1]))))
			tracker.GoodsDiff = append(tracker.GoodsDiff, data.CalulateGoodsDiff(allocation))

			// counter++
		}
	}

	tracker.EndTime = time.Now()
	fmt.Printf("Average Scores Diff: %.2f, Time elapsed: %v, Average Goods Diff: %.2f\n", tracker.AverageScoreDiff(), tracker.RunTime(), tracker.AverageGoodsDiff())
	fmt.Println()
}

func TestShallow(t *testing.T) {
	fmt.Println("===============Shallow Test Round Robin=================")
	dataPath := "../data"

	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// prepare for data
		if filepath.Ext(file.Name()) == ".txt" {
			parts := strings.Split(file.Name(), "_")
			if len(parts) >= 3 {
				if parts[1] != "shallow" {
					continue
				}

				n, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatal(err)
				}

				pattern := strings.TrimSuffix(parts[2], ".txt")
				if err != nil {
					log.Fatal(err)
				}

				fileData, err := data.ReadDataFile(filepath.Join(dataPath, file.Name()))
				if err != nil {
					log.Fatal(err)
				}

				dataFile := data.DataFile{
					N:         n,
					Pattern:   pattern,
					SampleNum: 1,
					FileName:  file.Name(),
					FileData:  fileData,
				}

				testRoundRobinShallow(dataFile)
			}
		}
	}
}

func testRoundRobinShallow(fileData data.DataFile) {
	tracker := data.Tracker{
		ScoresDiff: make([]int, 0),
		NumCases:   len(fileData.FileData),
	}

	fmt.Printf("Test for:\n\tN: %d, Pattern: %s, SampleNum: %d, FileName: %s, FileDataLen: %d\n", fileData.N, fileData.Pattern, fileData.SampleNum, fileData.FileName, len(fileData.FileData))

	cases := fileData.FileData
	N := fileData.N

	items := make([]*model.Item, N)
	for i := 0; i < N; i++ {
		items[i] = &model.Item{Name: "item" + strconv.Itoa(i+1)}
	}

	counter := 1
	tracker.StartTime = time.Now()

	// test process
	if fileData.Pattern == data.TIE {
		for i := 0; i < len(cases); i++ {

			preference1 := make(map[string]int)
			preference2 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference1[items[k].Name] = cases[i][k]
				preference2[items[k].Name] = cases[i][k]
			}

			fmt.Println("=======case " + strconv.Itoa(counter) + " ==========")
			fmt.Println(preference1)
			fmt.Println(preference2)

			agents := []*model.Agent{
				{"Alice", preference1, make(map[string]*model.Item)},
				{"Bob", preference2, make(map[string]*model.Item)},
			}

			allocation := roundRobin(agents, items)

			fmt.Println(allocation)
			score := data.CalculateScore(agents, allocation)

			fmt.Println("Scores:")
			values := []int{}
			for agentName, s := range score {
				values = append(values, s)
				fmt.Printf("%s: %d\n", agentName, s)
			}

			tracker.ScoresDiff = append(tracker.ScoresDiff, int(math.Abs(float64(values[0]-values[1]))))
			tracker.GoodsDiff = append(tracker.GoodsDiff, data.CalulateGoodsDiff(allocation))

			// counter++

		}
	} else {
		for i, j := 0, 1; i < len(cases); i, j = i+2, j+2 {

			preference1 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference1[items[k].Name] = cases[i][k]
			}

			preference2 := make(map[string]int)
			for k := 0; k < len(items); k++ {
				preference2[items[k].Name] = cases[j][k]
			}

			fmt.Println("=======case " + strconv.Itoa(counter) + " ==========")
			fmt.Println(preference1)
			fmt.Println(preference2)

			agents := []*model.Agent{
				{"Alice", preference1, make(map[string]*model.Item)},
				{"Bob", preference2, make(map[string]*model.Item)},
			}

			allocation := roundRobin(agents, items)

			fmt.Println(allocation)
			score := data.CalculateScore(agents, allocation)

			fmt.Println("Scores:")
			values := []int{}
			for agentName, s := range score {
				values = append(values, s)
				fmt.Printf("%s: %d\n", agentName, s)
			}

			tracker.ScoresDiff = append(tracker.ScoresDiff, int(math.Abs(float64(values[0]-values[1]))))
			tracker.GoodsDiff = append(tracker.GoodsDiff, data.CalulateGoodsDiff(allocation))

			// counter++
		}
	}

	tracker.EndTime = time.Now()
	fmt.Printf("Average Scores Diff: %.2f, Time elapsed: %v, Average Goods Diff: %.2f\n", tracker.AverageScoreDiff(), tracker.RunTime(), tracker.AverageGoodsDiff())
	fmt.Println()
}
