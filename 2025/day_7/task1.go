package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input.txt: %v", err)
	}
	defer f.Close()

	var grid []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n\r")
		if line != "" {
			grid = append(grid, line)
		} else {
			grid = append(grid, "")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	if len(grid) == 0 {
		log.Fatalf("input.txt is empty")
	}

	rows := len(grid)
	cols := len(grid[0])
	for i := range grid {
		if len(grid[i]) < cols {
			grid[i] = grid[i] + strings.Repeat(".", cols-len(grid[i]))
		}
	}

	startRow, startCol := -1, -1
	for r := 0; r < rows; r++ {
		if idx := strings.IndexRune(grid[r], 'S'); idx >= 0 {
			startRow = r
			startCol = idx
			break
		}
	}
	if startRow == -1 {
		log.Fatalf("no starting 'S' found in input")
	}

	activeCols := make(map[int]struct{})
	activeCols[startCol] = struct{}{}

	splits := 0

	for r := startRow + 1; r < rows; r++ {
		nextActive := make(map[int]struct{})
		for c := range activeCols {
			if c < 0 || c >= cols {
				continue
			}
			cell := grid[r][c]
			if cell == '^' {
				splits++
				if c-1 >= 0 {
					nextActive[c-1] = struct{}{}
				}
				if c+1 < cols {
					nextActive[c+1] = struct{}{}
				}
			} else {
				nextActive[c] = struct{}{}
			}
		}

		activeCols = nextActive
		if len(activeCols) == 0 {
			break
		}
	}

	fmt.Printf("Total splits: %d\n", splits)
}

