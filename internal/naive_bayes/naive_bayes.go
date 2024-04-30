package naive_bayes

import (
	"Naive_Bayes/internal/model"
	"math"
)

type ProbabilityMap struct {
	Probability      map[string]float64
	PriorProbability float64
	dataSize         int
}

func NewProbabilityMap(data *model.Data) *ProbabilityMap {
	result := ProbabilityMap{
		Probability: make(map[string]float64),
		dataSize:    data.CountOfFiles,
	}
	size := data.Size()
	for k, v := range data.WordCounts {
		result.Probability[k] = float64(v) / float64(size)
	}
	return &result
}

func NaiveBayes(normal, spam *ProbabilityMap, data model.Data) string {
	setPriorProbability(normal, spam)
	normalProbability := 1.0
	spamProbability := 1.0
	for k, v := range data.WordCounts {
		normalProbability *= math.Pow(normal.Probability[k], float64(v))
		spamProbability *= math.Pow(spam.Probability[k], float64(v))
	}
	if spamProbability > normalProbability {
		return "SPAM"
	} else {
		return "NORMAL"
	}
}

func setPriorProbability(normal *ProbabilityMap, spam *ProbabilityMap) {
	allDataSize := float64(spam.dataSize + normal.dataSize)
	normal.PriorProbability = float64(normal.dataSize) / allDataSize
	spam.PriorProbability = float64(spam.dataSize) / allDataSize
}
