package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/oowhyy/squaredle-solver/internal/words"
)

type Solver struct {
	wordList map[string]bool
	dirs     [][]int
}

func NewSolver() *Solver {
	wrds, err := words.Load()
	if err != nil {
		log.Fatal("unable to load words", err)
	}
	return &Solver{
		wordList: wrds,
		dirs:     [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
	}
}

func main() {
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
	solver := NewSolver()
	res := map[string]bool{}
	// solve
	solver.Solve(grid, res)
	// sort
	resList := make([]string, 0)
	for w := range res {
		resList = append(resList, w)
	}
	sort.Strings(resList)
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

func (s *Solver) Solve(grid [][]byte, res map[string]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			s.solveOne(grid, j, i, []byte{}, res)
		}
	}
}

func (s *Solver) solveOne(grid [][]byte, x, y int, pref []byte, res map[string]bool) {
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0 || grid[y][x] == '*' {
		return
	}
	curWord := append(pref, grid[y][x])
	if len(curWord) >= 4 && s.wordList[string(curWord)] {
		res[string(curWord)] = true
	}
	repair := grid[y][x]
	grid[y][x] = '*'
	for _, dir := range s.dirs {
		s.solveOne(grid, x+dir[0], y+dir[1], curWord, res)
	}
	grid[y][x] = repair
}
