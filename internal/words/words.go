package words

import (
	"bufio"
	"os"
	"strings"
)

func Load() (map[string]bool, error) {
	words := map[string]bool{}
	file, err := os.Open("wordList.txt")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 0 || word[0] == '#' {
			continue
		}
		words[strings.ToLower(word)] = true
	}
	return words, nil
}
