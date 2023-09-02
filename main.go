package main

import (
	"fmt"
	"merge/interval"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // strip out program name
	if len(args) == 1 {
		args = strings.Split(args[0], " ")
	}

	var input []interval.Interval[float64] // float64 is the largest type even though this is just for demo usage
	for _, arg := range args {
		in, err := interval.Parse[float64](arg)
		if err != nil {
			fmt.Println(err)
			return
		}
		input = append(input, in)
	}
	result, err := interval.Merge[float64](input...)
	if err != nil {
		fmt.Println(err)
	}

	resultStr := fmt.Sprintf("%v", result)
	fmt.Println(resultStr[1 : len(resultStr)-1]) // strip out brackets from slice stringification
}
