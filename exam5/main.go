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

		results[i] = createRoute(warehouse, robotA, robotB)

		for _, row := range results[i] {
			for _, r := range row {
				fmt.Printf("%c", r)
			}
			fmt.Println()
		}
	}
}

func createRoute(warehouse [][]rune, robotA, robotB Position) []string {

	rows := len(warehouse)
	cols := len(warehouse[0])

	result := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]rune, cols)
		copy(result[i], warehouse[i])
	}

	endPositionA := Position{0, 0}
	endPositionB := Position{rows - 1, cols - 1}
	if robotA.Row > robotB.Row {
		endPositionA, endPositionB = endPositionB, endPositionA
	} else if robotA.Row == robotB.Row && robotA.Col > robotB.Col {
		endPositionA, endPositionB = endPositionB, endPositionA
	}

	findPath(result, robotA, endPositionA, 'a')
	findPath(result, robotB, endPositionB, 'b')

	output := make([]string, rows)
	for r := range result {
		output[r] = string(result[r])
	}

	return output
}

func findPath(grid [][]rune, start, end Position, mark rune) {

	current := Position{start.Row, start.Col}

	stopRow := false
	stopCol := false

	for current.Row != end.Row || current.Col != end.Col {

		directionRow := 1
		directionCol := 1

		if current.Row > end.Row {
			directionRow = -1
		}
		if current.Col > end.Col {
			directionCol = -1
		}

		stopRow = current.Row == end.Row
		stopCol = current.Col == end.Col

		nextRow := current.Row + directionRow
		if !stopRow && grid[nextRow][current.Col] == '.' {
			grid[nextRow][current.Col] = mark
			current.Row = nextRow
			stopRow = current.Row == end.Row
			continue
		}
		nextCol := current.Col + directionCol
		if !stopCol && grid[current.Row][nextCol] == '.' {
			grid[current.Row][nextCol] = mark
			current.Col = nextCol
			stopCol = current.Col == end.Col
		}
	}
}
