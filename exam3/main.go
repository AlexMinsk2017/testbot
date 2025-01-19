package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {

		var n int
		fmt.Fscanf(reader, "%d\n", &n)

		var builder strings.Builder
		for j := 0; j < n; j++ {
			line, _ := reader.ReadString('\n')
			builder.WriteString(strings.TrimSpace(line))
		}
		var root Folder
		if err := json.Unmarshal([]byte(builder.String()), &root); err != nil {
			panic(err)
		}

		result := countBadFiles(&root)
		fmt.Println(result)
	}
}

func countBadFiles(folder *Folder) int {

	infected := 0
	findVirus := false

	for _, file := range folder.Files {
		if strings.HasSuffix(file, ".hack") {
			findVirus = true
			break
		}
	}

	if findVirus {
		return countAllFiles(folder)
	}

	for i := range folder.Folders {
		infected += countBadFiles(&folder.Folders[i])
	}

	return infected
}

func countAllFiles(folder *Folder) int {
	total := len(folder.Files)
	for i := range folder.Folders {
		total += countAllFiles(&folder.Folders[i])
	}
	return total
}
