package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Edge struct {
	weight float64
	v1     *Vertex
	v2     *Vertex
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) V1() *Vertex {
	return e.v1
}

func (e *Edge) V2() *Vertex {
	return e.v2
}

type Vertex struct {
	edges []*Edge
	num   int
}

func (v *Vertex) Num() int {
	return v.num
}

func (v *Vertex) Adj() []*Edge {
	return v.edges
}

func (v *Vertex) AddEdge(e *Edge) {
	v.edges = append(v.edges, e)
}

type Graph struct {
	vertices     []*Vertex
	num_vertices int
}

func (g *Graph) addVertices() {
	i := 0

	for i < g.num_vertices {
		v := new(Vertex)
		v.num = i
		g.vertices = append(g.vertices, v)
		i++
	}
}

func (g *Graph) Init(s string) *Graph {

	dat, err := ioutil.ReadFile(s)
	check(err)
	data := strings.Split(string(dat), "\n")

	nv, err := strconv.Atoi(data[0])
	g.num_vertices = nv
	g.addVertices()

	data = data[1:]
	for _, element := range data {
		elem_arr := strings.Split(element, " ")
		v1, err := strconv.Atoi(elem_arr[0])
		check(err)
		v2, err := strconv.Atoi(elem_arr[1])
		check(err)
		weight, err := strconv.ParseFloat(elem_arr[2], 64)
		check(err)

		e1 := Edge{weight, g.vertices[v1], g.vertices[v2]}
		e2 := Edge{weight, g.vertices[v2], g.vertices[v1]}
		g.vertices[v1].AddEdge(&e1)
		g.vertices[v2].AddEdge(&e2)
	}

	return g
}

func main() {
	g := new(Graph)
	g.Init("test.dat")
	for _, element := range g.vertices[0].Adj() {
		fmt.Printf("%f %d %d\n", element.weight, element.v1.Num(), element.v2.Num())
	}
}
