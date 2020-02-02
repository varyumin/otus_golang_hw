package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDataToSliceString(t *testing.T) {
	testData := "a b c d f e"
	testDataExpectedResult := []string{"a", "b", "c", "d", "f", "e"}
	data := DataToSliceString(testData)
	if cmp.Equal(data, testDataExpectedResult) != true {
		t.Log(data)
		t.Log(testDataExpectedResult)
		t.Fail()
	}
	t.Log(cmp.Equal(data, testDataExpectedResult))
}
func TestListToByCounter(t *testing.T) {
	testData := []string{"a", "a", "a", "b", "b", "c", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	testDataExpectedResult := ByCounter{
		Words{"a", 3},
		Words{"b", 2},
		Words{"c", 1},
		Words{"x", 9},
	}
	data := ListToByCounter(testData)
	if cmp.Equal(data, testDataExpectedResult) != true {
		t.Log(data)
		t.Log(testDataExpectedResult)
		t.Fail()
	}
	t.Log(cmp.Equal(data, testDataExpectedResult))
}

func TestSortWords(t *testing.T) {
	testData := ByCounter{
		Words{"a", 3},
		Words{"c", 1},
		Words{"b", 2},
		Words{"x", 9},
	}
	testDataExpectedResult := ByCounter{
		Words{"x", 9},
		Words{"a", 3},
		Words{"b", 2},
		Words{"c", 1},
	}
	data := SortWords(testData)
	if cmp.Equal(data, testDataExpectedResult) != true {
		t.Log(data)
		t.Log(testDataExpectedResult)
		t.Fail()
	}
	t.Log(cmp.Equal(data, testDataExpectedResult))
}
