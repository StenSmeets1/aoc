package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}

	rows := len(grid)
	cols := len(grid[0])

	dirs := []struct{ r, c int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, /*self*/ {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalRemoved := 0

	for {
		toRemove := make([][2]int, 0)

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] != '@' {
					continue
				}
				count := 0
				for _, d := range dirs {
					nr, nc := r+d.r, c+d.c
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
						count++
					}
				}
				if count < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break 
		}

		for _, pos := range toRemove {
			r, c := pos[0], pos[1]
			grid[r][c] = '.' 
		}

		totalRemoved += len(toRemove)
	}

	fmt.Println(totalRemoved)
}

