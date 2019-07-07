package main

import (
	"fmt"
	// "math"
)

func main() {
	// "e##e#o##oyof##q"
	// "e##e#fq##o##oyof##q"
	// "xywrrmp", "xywrrmu#p"
	// "bxj##tw"
	// "bxj###tw"
	// 	"bxj##tw"
	// "bxo#j##tw"
	res := calcEquation([][]string{[]string{"a", "b"}, []string{"b", "c"}}, []float64{2.0, 3.0}, [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}})
	fmt.Println(res)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]string, 0)
	res := make([]float64, len(queries))
	dist := make(map[string]float64, 0)

	for i := 0; i < len(equations); i++ {
		unity(m, dist, values[i], equations[i][0], equations[i][1])
		// fmt.Println(m)
	}
	// fmt.Println(m)
	// fmt.Println(dist)
	for i := 0; i < len(queries); i++ {
		_, ok1 := m[queries[i][0]]
		_, ok2 := m[queries[i][1]]
		if !ok1 || !ok2 {
			res[i] = -1.0
			continue
		}
		// fmt.Println("***")
		// fmt.Println(v1)
		// fmt.Println(v2)
		if find(m, dist, queries[i][0]) == find(m, dist, queries[i][1]) {
			// fmt.Println("*****")
			// fmt.Println(m)
			// fmt.Println(m)
			// fmt.Println(dist)
			res[i] = float64(dist[queries[i][0]] / dist[queries[i][1]])
		}
	}
	return res
}

func add(m map[string]string, dist map[string]float64, s string) {
	_, ok := m[s]

	if !ok {
		m[s] = s
		dist[s] = 1.0
	}

	return
}

func unity(m map[string]string, dist map[string]float64, value float64, s1 string, s2 string) {
	add(m, dist, s1)
	add(m, dist, s2)
	p1 := find(m, dist, s1)
	p2 := find(m, dist, s2)
	m[p1] = p2
	dist[p1] = float64((dist[s2] * value) / dist[s1])
	return
}

func find(m map[string]string, dist map[string]float64, s string) string {
	v, ok := m[s]

	if !ok {
		m[s] = s
	}

	if s != v {

		parent := find(m, dist, v)
		// fmt.Println("*****")
		// fmt.Println(s)
		// fmt.Println(parent)
		m[s] = parent
		dist[s] = dist[s] * dist[v]
		// fmt.Println(m)
		// fmt.Println(dist)
	}

	return m[s]
}
