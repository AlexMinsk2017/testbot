package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	Row, Col int
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	results := make([][]string, t)
	for i := 0; i < t; i++ {

		var rows, cols int
		fmt.Fscanf(reader, "%d %d\n", &rows, &cols)

		warehouse := make([][]rune, rows)
		var robotA, robotB Position

		for r := 0; r < rows; r++ {
			line, _ := reader.ReadString('\n')
			warehouse[r] = []rune(line[:cols])

			for c, cell := range warehouse[r] {
				if cell == 'A' {
					robotA = Position{r, c}
				} else if cell == 'B' {
					robotB = Position{r, c}
				}
			}
		}
	}
}
