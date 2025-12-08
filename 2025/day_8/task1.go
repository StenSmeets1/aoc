package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type P struct {
	x, y, z int64
}

type Edge struct {
	a, b int
	d    int64
}

type UF struct {
	p, sz []int
}

func NewUF(n int) *UF {
	p := make([]int, n)
	sz := make([]int, n)
	for i := range p {
		p[i] = i
		sz[i] = 1
	}
	return &UF{p: p, sz: sz}
}

func (u *UF) Find(x int) int {
	for u.p[x] != x {
		u.p[x] = u.p[u.p[x]]
		x = u.p[x]
	}
	return x
}

func (u *UF) Union(a, b int) bool {
	ra := u.Find(a)
	rb := u.Find(b)
	if ra == rb {
		return false
	}
	if u.sz[ra] < u.sz[rb] {
		ra, rb = rb, ra
	}
	u.p[rb] = ra
	u.sz[ra] += u.sz[rb]
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var pts []P
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			log.Fatalf("bad line: %q", line)
		}
		x, _ := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		y, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		z, _ := strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 64)
		pts = append(pts, P{x, y, z})
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	n := len(pts)
	var edges []Edge
	edges = make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			dy := pts[i].y - pts[j].y
			dz := pts[i].z - pts[j].z
			d2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{i, j, d2})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		if edges[i].d == edges[j].d {
			if edges[i].a == edges[j].a {
				return edges[i].b < edges[j].b
			}
			return edges[i].a < edges[j].a
		}
		return edges[i].d < edges[j].d
	})

	limit := 1000
	if limit > len(edges) {
		limit = len(edges)
	}

	uf := NewUF(n)
	for i := 0; i < limit; i++ {
		uf.Union(edges[i].a, edges[i].b)
	}

	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		r := uf.Find(i)
		counts[r]++
	}
	sizes := make([]int, 0, len(counts))
	for _, v := range counts {
		sizes = append(sizes, v)
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	prod := big.NewInt(1)
	for i := 0; i < 3; i++ {
		if i < len(sizes) {
			prod.Mul(prod, big.NewInt(int64(sizes[i])))
		} else {
			prod.Mul(prod, big.NewInt(1))
		}
	}
	fmt.Println(prod.String())
}

