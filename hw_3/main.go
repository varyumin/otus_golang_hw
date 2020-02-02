package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var (
	filterChats = []string{",", ".", "...", "-", "—", "?", "!", "\"", "(", ")", "[", "]", ":", ";", "«", "'", "*", "»"}
	file        = flag.String("file", "", "Path to text file")
	line        = flag.String("line", "", "String with values")
	topN        = flag.Int("top", 10, "Number of top")
)

type Words struct {
	Word    string
	Counter int
}

type ByCounter []Words

func TopN(n int, b []Words) []Words {
	var c []Words
	for i := 0; i < n; i++ {
		c = append(c, b[i])
	}
	return c
}

func printTop(b []Words) {
	for k, v := range b {
		fmt.Printf("%v place is taken by the word: '%v' occurs %v times\n", k+1, v.Word, v.Counter)
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

func fileToString(filepath string) (string, error) {
	dataByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dataByte), nil
}
func applyFilter(data string) string {
	data = removeFilterCharts(data)
	data = removeLineBreak(data)
	data = removeTab(data)
	data = toLowerText(data)
	data = removeTooManySpace(data)
	return data
}

func DataToSliceString(data string) []string {
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

func ListToByCounter(dataSlice []string) ByCounter {
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
func SortWords(b ByCounter) ByCounter {
	sort.Sort(ByCounter(b))
	return b
}

func main() {
	flag.Parse()
	var data string
	var err error

	if *file != "" {
		data, err = fileToString(*file)
		if err != nil {
			log.Panic(err)
		}
	} else {
		if *line == "" {
			log.Panic("Empty line")
		}
		data = *line
	}

	dataFiltered := applyFilter(data)
	dataSlice := DataToSliceString(dataFiltered)
	words := ListToByCounter(dataSlice)
	printTop(TopN(*topN, SortWords(words)))
}
