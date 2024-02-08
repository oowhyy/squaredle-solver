package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/oowhyy/squaredle-solver/internal/solver"
)

func main() {
	grid := loadInput()
	solver := solver.NewSolver()
	// solve find
	res := map[string]bool{}
	solver.Find(grid, res)
	// sort
	resList := make([]string, 0)
	for w := range res {
		resList = append(resList, w)
	}
	sort.Strings(resList)
	sort.SliceStable(resList, func(i, j int) bool { return len(resList[i]) < len(resList[j]) })
	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for _, word := range resList {
		out.WriteString(word + "\n")
	}
	log.Println("done")
}

func loadInput() [][]byte {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Fields(string(file))
	if len(lines) == 0 {
		log.Fatal("empty input")
	}
	m := len(lines[0])
	for _, line := range lines {
		if len(line) != m {
			log.Fatal("lines must be all equal length")
		}
	}
	n := len(lines)
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = []byte(lines[i])
	}
	return grid
}
