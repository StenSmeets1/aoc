package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type P struct{ x, y, z int64 }

type Edge struct {
	a, b int
	d    int64
}

type UF struct {
	p, sz []int
	comp  int
}

func NewUF(n int) *UF {
	p := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		sz[i] = 1
	}
	return &UF{p: p, sz: sz, comp: n}
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
	u.comp--
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
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		z, _ := strconv.ParseInt(parts[2], 10, 64)
		pts = append(pts, P{x, y, z})
	}

	n := len(pts)
	edges := make([]Edge, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			dy := pts[i].y - pts[j].y
			dz := pts[i].z - pts[j].z
			d2 := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{i, j, d2})
		}
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i].d < edges[j].d })

	uf := NewUF(n)

	for _, e := range edges {
		if uf.Union(e.a, e.b) {
			if uf.comp == 1 {
				ax := pts[e.a].x
				bx := pts[e.b].x
				fmt.Println(ax * bx)
				return
			}
		}
	}
}

