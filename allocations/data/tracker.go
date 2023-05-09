package data

import (
	"time"
)

type Tracker struct {
	StartTime  time.Time
	EndTime    time.Time
	ScoresDiff []int
	GoodsDiff  []int
	NumCases   int
}

func (t Tracker) RunTime() time.Duration {
	return t.EndTime.Sub(t.StartTime)
}

func (t Tracker) AverageScoreDiff() float64 {
	if len(t.ScoresDiff) == 0 {
		return 0
	}

	sum := 0
	n := 0

	for _, diff := range t.ScoresDiff {
		sum += diff
		n++
	}

	return float64(sum) / float64(n)
}

func (t Tracker) AverageGoodsDiff() float64 {
	if len(t.GoodsDiff) == 0 {
		return 0
	}

	sum := 0
	n := 0

	for _, diff := range t.GoodsDiff {
		sum += diff
		n++
	}

	return float64(sum) / float64(n)
}
