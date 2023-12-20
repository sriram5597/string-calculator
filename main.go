package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/sriram5597/calculator/string_calculator"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer
	fmt.Println("Enter Expression (Provide empty line at the end of expression): ")
	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		buffer.WriteString(line + "\n")
	}
	result := string_calculator.Sum(buffer.String())
	fmt.Println("result: ", result)
}
