package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var filterChats = []string{",", ".", "...", "-", "—", "?", "!", "\"", "(", ")", "[", "]", ":", ";", "«", "'", "*", "»"}

type Words struct {
	Word    string
	Counter int
}

type ByCounter []Words

func (b ByCounter) TopN(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%v place is taken by the word: '%v' occurs %v times\n", i+1, b[i].Word, b[i].Counter)
	}
}

func (b ByCounter) Len() int           { return len(b) }
func (b ByCounter) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByCounter) Less(i, j int) bool { return b[i].Counter > b[j].Counter }

func removeFilterCharts(data string) string {
	for _, chart := range filterChats {
		data = strings.ReplaceAll(data, chart, "")
	}
	return data
}

func removeTooManySpace(data string) string {
	var beg, end int
	beg = len(data)

	for {
		data = strings.ReplaceAll(data, "  ", " ")
		end = len(data)
		if beg == end {
			break
		} else {
			beg = end
		}
	}
	return data
}

func removeLineBreak(data string) string {
	return strings.ReplaceAll(data, "\n", " ")
}
func removeTab(data string) string {
	return strings.ReplaceAll(data, "\t", " ")
}

func toLowerText(data string) string {
	return strings.ToLower(data)
}

func openAndGet(filepath string) (string, error) {
	dataByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	data := string(dataByte)
	data = removeFilterCharts(data)
	data = removeLineBreak(data)
	data = removeTab(data)
	data = toLowerText(data)
	data = removeTooManySpace(data)
	return data, nil
}

func dataToSliceString(data string) []string {
	dataList := strings.Split(data, " ")
	for k, v := range dataList {
		if len(v) == 0 {
			removeIndex(dataList, k)
		}
	}
	return dataList
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func listToByCounter(dataSlice []string) ByCounter {
	var words ByCounter
	m := make(map[string]int)
	for _, v := range dataSlice {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		words = append(words, Words{k, v})
	}
	return words
}
func main() {
	data, _ := openAndGet("voyna-i-mir-tom-1.txt")
	dataSlice := dataToSliceString(data)
	words := listToByCounter(dataSlice)
	sort.Sort(ByCounter(words))
	words.TopN(100)
}
