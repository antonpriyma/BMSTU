package main

import (
        "fmt"
	"github.com/skorobogatov/input"
	//"strconv"
	//"bufio"
	//"os"
	//"github.com/derekparker/delve/pkg/dwarf/reader"
)

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

func main() {

	var m, n, q1, buf int
	var buf1 string

	input.Scanf("%d\n%d\n%d", &n, &m, &q1)
	g := make([]Edge, n)
	var preorder []int
	used := make(map[int]bool, n*m)
	gResult := make([]Edge, n)
	//fmt.Println("point",n,m,q1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			g[i].a = i
		}
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				switch i {
				case 0:
					input.Scanf("%d", &buf)
					//fmt.Println(buf)
					g[j].b = append(g[j].b, buf)
					//fmt.Println(g[j].b)
				case 1:
					input.Scanf("%s", &buf1)
					g[j].c = append(g[j].c, buf1)
				}
			}
		}
	}
	/*
		for i:=0;i<n;i++{
			fmt.Print(g[i].a," ")
			for j:=0;j<m;j++{
				fmt.Print(g[i].b[j]," ")
				fmt.Print(g[i].c[j]," ")
			}
		//	fmt.Println()
		}*/

	//fmt.Println("point",n)
	//preorder=append(preorder,q1 )
	preorder = dfs(g, used, q1, preorder)
	//fmt.Println(preorder)

	q1 = 0
	for i := 0; i < len(preorder); i++ {
		gResult[i] = g[preorder[i]]
		gResult[i].a = i

		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				if preorder[k] == g[preorder[i]].b[j] {
					buf = k
					break
				}
			}
			gResult[i].b[j] = buf
			gResult[i].c[j] = g[preorder[i]].c[j]
		}
		//fmt.Println(gResult[i].a,gResult[i].b,gResult[i].c)
	}

	fmt.Printf("%d\n%d\n%d\n", len(preorder), m, q1)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(preorder); j++ {
			for k := 0; k < m; k++ {
				if k > 0 {
					fmt.Print(" ")
				}
				switch i {
				case 0:
					fmt.Print(gResult[j].b[k])
				case 1:
					fmt.Print(gResult[j].c[k])
				}
			}
			fmt.Println()
		}
	}

	/*
		fmt.Println("digraph {")
		fmt.Printf("\t%s\n","rankdir = LR")
		fmt.Printf("\t%s\n","dummy [label = \"\", shape = none]")
		for i:=0;i<m;i++{
			fmt.Printf("\t%d [shape = circle]\n",i)
		}
		fmt.Println("\tdummy ->",q1)
		for i:=0;i<m;i++{
			for j:=0;j<n;j++ {
				//fmt.Println(i,"->", delta[i][j]," [label = ",string(j+'a'),"(",dunc[i][j],")]")
				fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n",i,delta[i][j],string(j+'a'),dunc[i][j])
			}
		}
		fmt.Println("}")
	*/
}
