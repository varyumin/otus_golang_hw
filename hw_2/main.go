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

var LineWrap string

func StringToInt(numberStr string) (int, error) {
	if numberStr == "" {
		return 1, nil
	} else {
		number, err := strconv.ParseFloat(numberStr, 10)
		if err != nil {
			return 0, fmt.Errorf("Cannot convert int to string: %s", err)
		}
		return int(number), nil
	}
}
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
		count, err := StringToInt(v[2])
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
	flag.StringVar(&LineWrap, "string", "", "primitive unboxing of a string containing duplicate characters/runes")
	flagenv.Parse()
	flag.Parse()

	duplicate, err := Unpack(LineWrap)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(duplicate)
}
