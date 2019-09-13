package main

import (
	"fmt"
	"os"
	"strings"
)

func getSecond(arg []string, Value []string) {
	for _, value := range Value {
		if strings.EqualFold(arg[0], value) {
			// Match found
			fmt.Println("Value: ", value)
			break
		}
	}
}

func getFirst(args []string, m map[string][]string) {
	argLen := len(args) - 1

	// Loop through args
	for i := 1; i <= argLen; i += 2 {
		key := args[i]
		fmt.Println("\nKey: ", key)
		// Search in m for expected second argument
		value, ok := m[key]
		if !ok {
			fmt.Println("Argument unmatched")
			// os.Exit(0)
		}
		// Search in value for second argument
		fmt.Println(args[i+1:i+2], " ", i+1, " ", i+2, value)
		getSecond(args[i+1:i+2], value)

	}
}

func main() {
	m := map[string][]string{
		"-s": {"seeds", "peers", "rating"},
		"-v": {"seeds", "peers", "rating", "title", "quality"},
		"-h": {},
	}
	getFirst(os.Args, m)
}
