package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [][2]int{
	{0, 1},  // вправо
	{1, 0},  // вниз
	{1, 1},  // диагональ вниз-вправо
	{1, -1}, // диагональ вниз-влево
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	for kit := 0; kit < t; kit++ {

		var k, n, m int
		fmt.Fscan(reader, &k)
		fmt.Fscan(reader, &n, &m)

		reader.ReadString('\n')

		board := make([][]rune, n)
		for i := 0; i < n; i++ {
			line, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			board[i] = []rune(strings.TrimSpace(line))
		}
		if dfs(board, k, n, m) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func check(board [][]rune, k, n, m int, player rune) bool {

	inBorders := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != player {
				continue
			}
			for _, dir := range directions {
				count := 1
				x, y := i+dir[0], j+dir[1]
				for inBorders(x, y) && board[x][y] == player {
					count++
					if count == k {
						return true
					}
					x += dir[0]
					y += dir[1]
				}
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

func checkFromCell(board [][]rune, k, n, m, x, y int, player rune) bool {
	inBorders := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < m
	}
	for _, dir := range directions {
		count := 1
		dx, dy := dir[0], dir[1]

		nx, ny := x+dx, y+dy
		for inBorders(nx, ny) && board[nx][ny] == player {
			count++
			nx += dx
			ny += dy
		}

		nx, ny = x-dx, y-dy
		for inBorders(nx, ny) && board[nx][ny] == player {
			count++
			nx -= dx
			ny -= dy
		}

		if count >= k {
			return true
		}
	}
	return false
}
