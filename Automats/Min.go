package main

import (
        "github.com/skorobogatov/input"
	//"fmt"
	//"strconv"
	"fmt"
)

type DSU struct {
	parent []int
	rank   []int
}

type Mealy struct {
	fun   [][]string
	delta [][]int
	q     [][]int
	n     int
	m     int
	q1    int
}

type Edge struct {
	a int
	b []int
	c []string
}

type Edge1 struct {
	a int
	b int
	c string
}

func dfs(graph []Edge, used map[int]bool, node int, preorder []int) []int {
	preorder = append(preorder, graph[node].a)
	used[node] = true
	//fmt.Println("a: ",node,graph[node].a,graph[node].b,preorder)
	//fn(node)
	for _, i := range graph[node].b {
		//fmt.Println("b :",i,used[i])
		if _, ok := used[i]; !ok {
			preorder = dfs(graph, used, i, preorder)
		}
	}

	//fmt.Println()
	return preorder
}

func (avt *Mealy) canon() Mealy {
	var g []Edge
	g = avt.toGraph()
	var preorder []int
	used := make(map[int]bool, avt.n*avt.m)
	gResult := make([]Edge, avt.n)
	var avtResult Mealy
	avtResult.n = avt.n
	avtResult.m = avt.m
	avtResult.q1 = 0
	avtResult.delta = make([][]int, avt.n)
	avtResult.fun = make([][]string, avt.n)
	buf := 0
	preorder = dfs(g, used, avt.q1, preorder)
	for i := 0; i < len(preorder); i++ {
		gResult[i] = g[preorder[i]]
		gResult[i].a = i

		for j := 0; j < avt.m; j++ {
			for k := 0; k < avt.n; k++ {
				if preorder[k] == g[preorder[i]].b[j] {
					buf = k
					break
				}
			}
			gResult[i].b[j] = buf
			gResult[i].c[j] = g[preorder[i]].c[j]
		}
		avtResult.n = len(preorder)
		//fmt.Println(gResult[i].a,gResult[i].b,gResult[i].c)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < len(preorder); j++ {
			for k := 0; k < avt.m; k++ {
				if k > 0 {
					fmt.Print(" ")
				}
				switch i {
				case 0:
					avtResult.delta[j] = append(avtResult.delta[j], gResult[j].b[k])
				case 1:
					avtResult.fun[j] = append(avtResult.fun[j], gResult[j].c[k])
				}
			}
			fmt.Println()
		}
	}
	return avtResult
}

func (avt *Mealy) toGraph() []Edge {
	g := make([]Edge, avt.n)
	for i := 0; i < avt.n; i++ {
		for j := 0; j < avt.m; j++ {
			g[i].a = i
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < avt.n; j++ {
			for k := 0; k < avt.m; k++ {
				switch i {
				case 0:
					//fmt.Println(buf)
					g[j].b = append(g[j].b, avt.delta[j][k])
					//fmt.Println(g[j].b)
				case 1:
					g[j].c = append(g[j].c, avt.fun[j][k])
				}
			}
		}
	}
	return g
}

func (q *DSU) make_set(v int) {
	q.parent[v] = v
	q.rank[v] = 0
}

func (q *DSU) find_set(v int) int {
	if v == q.parent[v] {
		return v
	}
	q.parent[v] = q.find_set(q.parent[v])
	return q.parent[v]
}

func (q *DSU) union_sets(a int, b int) {
	a = q.find_set(a)
	b = q.find_set(b)
	if a != b {
		if q.rank[a] < q.rank[b] {
			c := a
			a = b
			b = c
		}
		q.parent[b] = a
		if q.rank[a] == q.rank[b] {
			q.rank[a]++
		}
	}
}

func (avt *Mealy) split1(pi []int) (int, []int) {
	m := avt.n
	var s DSU
	s.parent = make([]int, avt.n)
	s.rank = make([]int, avt.n)
	for i := 0; i < avt.n; i++ {
		s.make_set(i)
	}
	for i := 0; i < avt.n; i++ {
		for j := i + 1; j < avt.n; j++ {
			if s.find_set(i) != s.find_set(j) {
				eq := true
				for k := 0; k < avt.m; k++ {
					if avt.fun[i][k] != avt.fun[j][k] {
						eq = false
						break
					}
				}
				if eq {
					s.union_sets(i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < avt.n; i++ {
		pi[i] = s.find_set(i)
	}
	return m, pi
}

func (avt *Mealy) split(pi []int) (int, []int) {
	m := avt.n
	var s DSU
	s.parent = make([]int, avt.n)
	s.rank = make([]int, avt.n)
	for i := 0; i < avt.n; i++ {
		s.make_set(i)
	}

	for i := 0; i < avt.n; i++ {
		for j := i + 1; j < avt.n; j++ {
			if pi[i] == pi[j] && s.find_set(i) != s.find_set(j) {
				eq := true
				for k := 0; k < avt.m; k++ {
					w1 := avt.delta[i][k]
					w2 := avt.delta[j][k]
					if pi[w1] != pi[w2] {
						eq = false
						break
					}
				}
				if eq {
					s.union_sets(i, j)
					m--
				}
			}
		}
	}
	for i := 0; i < avt.n; i++ {
		pi[i] = s.find_set(i)
	}
	return m, pi
}

func (avt *Mealy) AufenkampHohn() Mealy {
	pi := make([]int, avt.n)
	m := -1
	m_ := -1
	m, pi = avt.split1(pi)
	//fmt.Println(pi)
	for {
		if m == m_ {
			break
		}
		m_ = m
		m, pi = avt.split(pi)
		//fmt.Println(pi)
	}
	var buf Mealy
	buf.n = m
	buf.m = avt.m
	buf.q1 = 0
	buf.delta = make([][]int, m)
	buf.fun = make([][]string, m)
	q_ := 0
	for j := 0; j < m; j++ {
		buf.delta[j] = make([]int, avt.m)
		buf.fun[j] = make([]string, avt.m)
	}
	vertex := make([]int, avt.n)
	vertexResult := make([]int, avt.n)

	for i := 0; i < avt.n; i++ {
		if pi[i] == i {
			vertex[q_] = i
			vertexResult[i] = q_
			q_++
		}
	}

	buf.q1 = vertexResult[pi[avt.q1]]
	for i := 0; i < buf.n; i++ {
		for j := 0; j < buf.m; j++ {
			buf.delta[i][j] = vertexResult[pi[avt.delta[vertex[i]][j]]]
			buf.fun[i][j] = avt.fun[vertex[i]][j]
		}
	}
	return buf
}

func (avt *Mealy) print() {
	fmt.Printf("%d\n%d\n%d\n", avt.n, avt.m, avt.q1)
	for i := 0; i < 2; i++ {
		for j := 0; j < avt.n; j++ {
			for k := 0; k < avt.m; k++ {
				switch i {
				case 0:
					fmt.Print(avt.delta[j][k], " ")
				case 1:
					fmt.Print(avt.fun[j][k], " ")
				}
			}
			fmt.Println()
		}
	}

}

func (avt *Mealy) printLikeDOT() {
	fmt.Println("digraph {")
	fmt.Printf("\t%s\n", "rankdir = LR")
	fmt.Printf("\t%s\n", "dummy [label = \"\", shape = none]")
	for i := 0; i < avt.n; i++ {
		fmt.Printf("\t%d [shape = circle]\n", i)
	}
	fmt.Println("\tdummy ->", avt.q1)
	for i := 0; i < avt.n; i++ {
		for j := 0; j < avt.m; j++ {
			//fmt.Println(i,"->", delta[i][j]," [label = ",string(j+'a'),"(",dunc[i][j],")]")
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i, avt.delta[i][j], string(j+'a'), avt.fun[i][j])
		}
	}
	fmt.Println("}")

}

func main() {

	var m, n, q1 int

	input.Scanf("%d\n%d\n%d", &n, &m, &q1)
	var avt Mealy
	avt.delta = make([][]int, n)
	avt.fun = make([][]string, n)
	avt.n = n
	avt.m = m
	avt.q1 = q1
	for j := 0; j < n; j++ {
		avt.delta[j] = make([]int, m)
		avt.fun[j] = make([]string, m)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				switch i {
				case 0:
					input.Scanf("%d", &avt.delta[j][k])
				case 1:
					input.Scanf("%s", &avt.fun[j][k])
				}
			}
		}
	}

	result := avt.AufenkampHohn()
	avtResult := result.canon()
	avtResult.printLikeDOT()

}
