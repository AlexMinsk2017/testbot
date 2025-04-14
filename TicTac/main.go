package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [][2]int{
	{0, 1}, {1, 0},
	{1, 1}, {1, -1},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for test := 0; test < t; test++ {
		var k, n, m int
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n, &m)

		board := make([][]rune, n)
		for i := 0; i < n; i++ {
			line, _ := reader.ReadString('\n')
			board[i] = []rune(strings.TrimSpace(line))
		}

		if dfs(board, k, n, m) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func inBorders(x, y, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func checkFromCell(board [][]rune, k, n, m, x, y int, player rune) bool {
	for _, dir := range directions {
		count := 1
		for step := 1; step < k; step++ {
			nx, ny := x+dir[0]*step, y+dir[1]*step
			if inBorders(nx, ny, n, m) && board[nx][ny] == player {
				count++
			} else {
				break
			}
		}
		for step := 1; step < k; step++ {
			nx, ny := x-dir[0]*step, y-dir[1]*step
			if inBorders(nx, ny, n, m) && board[nx][ny] == player {
				count++
			} else {
				break
			}
		}
		if count >= k {
			return true
		}
	}
	return false
}

func check(board [][]rune, k, n, m int, player rune) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != player {
				continue
			}
			if checkFromCell(board, k, n, m, i, j, player) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]rune, k, n, m int) bool {
	if check(board, k, n, m, 'X') || check(board, k, n, m, 'O') {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != '.' {
				continue
			}
			board[i][j] = 'X'
			if checkFromCell(board, k, n, m, i, j, 'X') {
				board[i][j] = '.'
				return true
			}
			board[i][j] = '.'
		}
	}
	return false
}
