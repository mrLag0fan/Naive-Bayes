package scanner

import (
	"Naive_Bayes/internal/model"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ScanTrainingDataInFolder(folderPath string) *model.Data {
	var allDataFromFolder *model.Data = &model.Data{
		WordCounts:   make(map[string]int),
		CountOfFiles: 0,
	}
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			fmt.Printf("Scanning file: %s\n", path)
			wordCounts, err := scanData(path)
			if err != nil {
				fmt.Printf("Error scanning file %s: %s\n", path, err)
				return nil
			}
			allDataFromFolder.Add(wordCounts)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning folder %s: %s\n", folderPath, err)
	}
	return allDataFromFolder
}

func scanData(path string) (*model.Data, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	res := model.Data{
		WordCounts: make(map[string]int),
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		for _, word := range line {
			res.WordCounts[word]++
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return &res, nil
}

func AddAlfa(normal, spam *model.Data) {
	for word := range spam.WordCounts {
		if _, ok := normal.WordCounts[word]; !ok {
			normal.WordCounts[word] = 1
		} else {
			normal.WordCounts[word] += 1
		}
	}

	// Add words from normal to spam with a value of 1 if they don't already exist
	for word := range normal.WordCounts {
		if _, ok := spam.WordCounts[word]; !ok {
			spam.WordCounts[word] = 1
		} else {
			spam.WordCounts[word] += 1
		}
	}
}

func ScanFilesInFolder(folderPath string) []*model.Data {
	var allDataFromFolder []*model.Data

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			fmt.Printf("Scanning file: %s\n", path)
			wordCounts, err := scanData(path)
			if err != nil {
				fmt.Printf("Error scanning file %s: %s\n", path, err)
				return nil
			}
			allDataFromFolder = append(allDataFromFolder, wordCounts)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error scanning folder %s: %s\n", folderPath, err)
	}

	return allDataFromFolder
}
