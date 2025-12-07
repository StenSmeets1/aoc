package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var grid []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) == 0 {
		log.Fatal("input is empty after filtering")
	}

	rows := len(grid)
	cols := len(grid[0])

	for i, row := range grid {
		if len(row) != cols {
			log.Fatalf("row %d has length %d but expected %d", i, len(row), cols)
		}
	}

	sr, sc := -1, -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				sr = r
				sc = c
				break
			}
		}
		if sr != -1 {
			break
		}
	}
	if sr == -1 {
		log.Fatal("No S found in the grid")
	}

	curr := make([]*big.Int, cols)
	next := make([]*big.Int, cols)
	for i := range curr {
		curr[i] = big.NewInt(0)
		next[i] = big.NewInt(0)
	}

	curr[sc].SetInt64(1) 

	for r := sr; r+1 < rows; r++ {
		for i := 0; i < cols; i++ {
			next[i].SetInt64(0)
		}

		for c := 0; c < cols; c++ {
			if curr[c].Sign() == 0 {
				continue
			}

			down := grid[r+1][c]

			if down == '^' {
				if c-1 >= 0 {
					next[c-1].Add(next[c-1], curr[c])
				}
				if c+1 < cols {
					next[c+1].Add(next[c+1], curr[c])
				}
			} else {
				next[c].Add(next[c], curr[c])
			}
		}

		curr, next = next, curr
	}

	total := big.NewInt(0)
	for _, x := range curr {
		total.Add(total, x)
	}

	fmt.Println("Total timelines:", total.String())
}

