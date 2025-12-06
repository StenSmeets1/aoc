package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := []string{}
	maxLen := 0
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	for i := range lines {
		if len(lines[i]) < maxLen {
			lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i]))
		}
	}

	hasNonSpace := make([]bool, maxLen)
	for col := 0; col < maxLen; col++ {
		for _, line := range lines {
			if line[col] != ' ' {
				hasNonSpace[col] = true
				break
			}
		}
	}

	spans := [][2]int{}
	inSpan := false
	start := 0
	for col := 0; col < maxLen; col++ {
		if hasNonSpace[col] {
			if !inSpan {
				inSpan = true
				start = col
			}
		} else {
			if inSpan {
				inSpan = false
				spans = append(spans, [2]int{start, col - 1})
			}
		}
	}
	if inSpan {
		spans = append(spans, [2]int{start, maxLen - 1})
	}

	grandTotal := big.NewInt(0)

	for _, sp := range spans {
		l, r := sp[0], sp[1]
		tokens := []string{}
		for _, line := range lines {
			segment := strings.TrimSpace(line[l : r+1])
			if segment != "" {
				tokens = append(tokens, segment)
			}
		}
		if len(tokens) == 0 {
			continue
		}

		op := tokens[len(tokens)-1]   
		numTokens := tokens[:len(tokens)-1] 

		result := big.NewInt(0)
		if op == "+" {
			for _, t := range numTokens {
				n := new(big.Int)
				_, ok := n.SetString(t, 10)
				if !ok {
					panic("failed to parse integer: " + t)
				}
				result.Add(result, n)
			}
		} else if op == "*" {
			result.SetInt64(1)
			for _, t := range numTokens {
				n := new(big.Int)
				_, ok := n.SetString(t, 10)
				if !ok {
					panic("failed to parse integer: " + t)
				}
				result.Mul(result, n)
			}
		} else {
			panic("unknown operator: " + op)
		}

		grandTotal.Add(grandTotal, result)
	}

	fmt.Println(grandTotal.String())
}

