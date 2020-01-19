package main

import (
	"flag"
	"fmt"
	"github.com/facebookgo/flagenv"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lineWrap string

func stringToInt(numberStr string) (int, error) {
	if numberStr == "" {
		return 1, nil
	}
	number, err := strconv.ParseFloat(numberStr, 10)
	if err != nil {
		return 0, fmt.Errorf("Cannot convert int to string: %s", err)
	}
	return int(number), nil
}

//Unpack unboxing of a string containing duplicate characters/runes
func Unpack(wrapped string) (string, error) {
	if len(wrapped) == 0 {
		return "", nil
	}
	var unwrapped string
	re, err := regexp.Compile("([a-zA-Z])([0-9]*)")
	if err != nil {
		return "", err
	}
	res := re.FindAllStringSubmatch(wrapped, -1)
	for _, v := range res {
		count, err := stringToInt(v[2])
		if err != nil {
			return "", err
		}
		unwrapped = unwrapped + strings.Repeat(v[1], count)
	}
	if len(unwrapped) == 0 {
		return "", fmt.Errorf("Invalid string")
	}
	return unwrapped, nil
}

func main() {
	flag.StringVar(&lineWrap, "line", "", "primitive unboxing of a string containing duplicate characters/runes")
	flagenv.Parse()
	flag.Parse()

	duplicate, err := Unpack(lineWrap)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(duplicate)
}
