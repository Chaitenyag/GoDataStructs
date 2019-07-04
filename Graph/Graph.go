package main

import (
	"container/list"
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
	edges list.List
	num   int32
}

func (v *Vertex) Num() int32 {
	return v.num
}

func (v *Vertex) Adj() list.List {
	return v.edges
}

func (v *Vertex) AddEdge(e *Edge) {
	v.edges.PushBack(e)
}

type Graph struct {
	vertices     list.List
	num_vertices int32
}

func graph(s string) *Graph {
	g := new(Graph)
	dat, err := ioutil.ReadFile(s)
	check(err)
	data := strings.Split(string(dat), "\n")
	nv, err := strconv.Atoi(data[0])
	data = data[1:]
	for _, element := range data {
		elem_arr := strings.Split(element, " ")
		v1, err := strconv.Atoi(elem_arr[0])
		v2, err := strconv.Atoi(elem_arr[1])
		g.vertices[int32(v1)].AddEdge()
	}
	fmt.Println(data)

	g.num_vertices = int32(nv)

	return g
}

func main() {
	g := graph("test.dat")
	fmt.Println(g)
}
