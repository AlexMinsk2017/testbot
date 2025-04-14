package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTicTac(t *testing.T) {
	// Путь к папке с тестами
	dir := "D:/Work/OZON/bot/TicTac/test"

	// Количество тестов
	const totalTests = 38

	for i := 1; i <= totalTests; i++ {
		// Составляем пути к файлам
		inputPath := filepath.Join(dir, fmt.Sprintf("%d", i))
		answerPath := filepath.Join(dir, fmt.Sprintf("%d.a", i))

		// Проверяем, существует ли файл ввода и ответа
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			t.Logf("Пропущен тест %d: файл ввода не найден", i)
			continue
		}
		if _, err := os.Stat(answerPath); os.IsNotExist(err) {
			t.Logf("Пропущен тест %d: файл ответа не найден", i)
			continue
		}

		// Используем t.Run для отдельного отображения каждого теста
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			inputFile, err := os.Open(inputPath)
			if err != nil {
				t.Fatalf("Ошибка открытия файла ввода %s: %v", inputPath, err)
			}
			defer inputFile.Close()

			answerFile, err := os.Open(answerPath)
			if err != nil {
				t.Fatalf("Ошибка открытия файла ответа %s: %v", answerPath, err)
			}
			defer answerFile.Close()

			inputReader := bufio.NewReader(inputFile)
			answerReader := bufio.NewReader(answerFile)

			var numTests int
			fmt.Fscan(inputReader, &numTests)

			for kit := 0; kit < numTests; kit++ {
				var k, n, m int
				fmt.Fscan(inputReader, &k)
				fmt.Fscan(inputReader, &n, &m)
				inputReader.ReadString('\n') // пропускаем оставшуюся часть строки

				board := make([][]rune, n)
				for j := 0; j < n; j++ {
					line, err := inputReader.ReadString('\n')
					if err != nil {
						t.Fatalf("Ошибка чтения строки доски из файла %d: %v", i, err)
					}
					board[j] = []rune(strings.TrimSpace(line))
				}

				var result string
				if dfs(board, k, n, m) {
					result = "YES"
				} else {
					result = "NO"
				}

				expected, err := answerReader.ReadString('\n')
				if err != nil {
					t.Fatalf("Ошибка чтения строки ответа из файла %d.a: %v", i, err)
				}
				expected = strings.TrimSpace(expected)

				if result != expected {
					t.Errorf("Кейс #%d: ожидалось %s, получено %s", kit+1, expected, result)
				} else {
					fmt.Print(result + " - " + expected + "\n")
				}
			}
		})
	}
}
