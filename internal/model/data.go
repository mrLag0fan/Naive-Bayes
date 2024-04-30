package model

type Data struct {
	WordCounts   map[string]int
	CountOfFiles int
}

func (data *Data) Size() int {
	size := 0
	for _, v := range data.WordCounts {
		size += v
	}
	return size
}

func (data *Data) Add(other *Data) {
	for k, v := range other.WordCounts {
		data.WordCounts[k] += v
	}
	data.CountOfFiles++
}
