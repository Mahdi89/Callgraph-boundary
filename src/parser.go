//DOT Graph analyzer
//Author: Mahdi Jelodari
//Date: April 2018
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	name string
	next *Node
}

type Node_ struct {
	name string
	id   string
}

func Map(l_ *list.List, caller string, callee string) (string, string) {

	var caller_id, callee_id string

	for e := l_.Front(); e != nil; e = e.Next() {
		if e.Value.(Node_).name == caller {
			caller_id = e.Value.(Node_).id
		}
		if e.Value.(Node_).name == callee {
			callee_id = e.Value.(Node_).id
		}
	}

	return caller_id, callee_id
}

func Parse(file string) (*list.List, *list.List) {

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	// Parse the generated dot graph and return a list
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	l := list.New()
	l_ := list.New()

	for scanner.Scan() {
		str := scanner.Text()
		if s.HasPrefix(str, "Node") {
			scanner.Scan()
			nxt := scanner.Text()
			if s.HasPrefix(nxt, "->") {
				scanner.Scan()
				str2 := scanner.Text()
				e1 := Node{str2, nil}
				e2 := Node{str, &e1}
				l.PushBack(e2)

			} else {
				e1 := Node_{nxt, str}
				l_.PushBack(e1)
			}
		}
	}
	return l, l_
}

// Identical to list.find (by id or name)
func Existby(l *list.List, call string, id_name bool) bool {

	if id_name {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(Node_).id == call {
				return true
			}
		}
		return false
	}else {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(Node).name == call {
				return true
			}
		}
		return false
	}
}


func Search(l *list.List, visited *list.List, caller string, callee string) bool {

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(Node).name == caller {
			nxt := e.Value.(Node).next
			visited.PushBack(e.Value.(Node))
			nm := nxt.name
			if nm == callee{
				return true
			} else if !Existby(visited, nm, false){
				if Search(l, visited, nm, callee) {
					return true
				} else {
					continue
				}
			}
		}
	}
	return false
}

func main() {

	file := "./callgraph.dot"

	// Parse the input file
	l, l_ := Parse(file)

	args := os.Args[1:]
	caller := args[0]
	callee := args[1]

	// Map names to ids
	caller, callee = Map(l_, caller, callee)

	// Keep a list of visited calls to avoid cycles
	visited := list.New()	

	if Existby(l_, caller, true) && Existby(l_, callee, true) {
		// TODO use l_ to map caller/callee to actual names
		fmt.Println(Search(l, visited, caller, callee))
	} else {
		fmt.Println("Please check: Callee/Caller or both dosn't exist!")
	}

}
