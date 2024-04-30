package main

import (
	"Naive_Bayes/internal/naive_bayes"
	"Naive_Bayes/internal/scanner"
	"fmt"
)

func main() {
	normalTrainingData := scanner.ScanTrainingDataInFolder("./data/normal")
	spamTrainingData := scanner.ScanTrainingDataInFolder("./data/spam")
	unknownData := scanner.ScanFilesInFolder("./data/unknown")
	scanner.AddAlfa(normalTrainingData, spamTrainingData)
	fmt.Println(normalTrainingData)
	fmt.Println(spamTrainingData)
	normalTrainedProbability := naive_bayes.NewProbabilityMap(normalTrainingData)
	spamTrainedProbability := naive_bayes.NewProbabilityMap(spamTrainingData)
	fmt.Println(normalTrainedProbability)
	fmt.Println(spamTrainedProbability)
	bayesNormal := naive_bayes.NaiveBayes(normalTrainedProbability, spamTrainedProbability, *normalTrainingData)
	bayesSpam := naive_bayes.NaiveBayes(normalTrainedProbability, spamTrainedProbability, *spamTrainingData)
	fmt.Println(bayesNormal)
	fmt.Println(bayesSpam)

	for _, data := range unknownData {
		bayes := naive_bayes.NaiveBayes(normalTrainedProbability, spamTrainedProbability, *data)
		fmt.Printf("%v : %s\n", *data, bayes)
	}
}
