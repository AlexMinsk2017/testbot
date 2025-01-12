package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &data[i])
	}

	salaries := make([]int, n)
	for i, value := range data {
		valueStr := strconv.Itoa(value)
		if len(valueStr) == 1 {
			salaries[i] = 0
			continue
		}

		salary := 0
		for j := 0; j < len(valueStr); j++ {
			newStr := valueStr[:j] + valueStr[j+1:]
			newSalary, _ := strconv.Atoi(newStr)
			salary = max(salary, newSalary)
		}
		salaries[i] = salary
	}

	for _, salary := range salaries {
		fmt.Fprintf(out, "%d\n", salary)
	}
}
