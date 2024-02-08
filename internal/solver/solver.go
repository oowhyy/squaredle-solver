package solver

import (
	"log"

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

// try all possible paths on grid, very slow for big grids
func (s *Solver) Solve(grid [][]byte, res map[string]bool) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			s.solveOne(grid, j, i, []byte{}, res)
		}
	}
}

func (s *Solver) solveOne(grid [][]byte, x, y int, pref []byte, res map[string]bool) {
	if x >= len(grid[0]) || y >= len(grid) || x < 0 || y < 0 || grid[y][x] == '*' {
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

// find words from wordList in grid
// find word for each word in dictionary - works FASTER than solve
func (s *Solver) Find(grid [][]byte, res map[string]bool) {
	for w := range s.wordList {
		if len(w) < 4 {
			continue
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				s.tryFindOne(grid, j, i, []byte{}, w, res)
			}
		}
	}
}

func (s *Solver) tryFindOne(grid [][]byte, x, y int, pref []byte, target string, res map[string]bool) {
	if target == "" {
		return
	}
	if x >= len(grid[0]) || y >= len(grid) || x < 0 || y < 0 {
		return
	}
	if grid[y][x] != target[len(pref)] {
		return
	}
	curWord := append(pref, grid[y][x])
	if string(curWord) == target {
		res[target] = true
		return
	}

	repair := grid[y][x]
	grid[y][x] = '*'
	for _, dir := range s.dirs {
		s.tryFindOne(grid, x+dir[0], y+dir[1], curWord, target, res)
	}
	grid[y][x] = repair
}
