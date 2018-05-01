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

func Parse(file string) (*list.List, *list.List)  {

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

			}else {
				e1 := Node_{nxt, str}
				l_.PushBack(e1)
			}
		}
	}
	return l,l_
}

func Search(l *list.List, caller string, callee string) bool {


	tf := make(chan bool, 2)
	go func() {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(Node).name == caller {
				tf <- true
				break
			}
		}
		tf <- false
	}()

	go func() {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value.(Node).name == callee {
				tf <- true
				break
			}
		}
		tf <- false
	}()

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(Node).name == caller {
			nxt := e.Value.(Node).next
			nm := nxt.name
			fmt.Println(nm)
			if nm == callee {
				return true
			} else {
				if Search(l, nm, callee) { 
					return true
				} else  {continue}
			}
		}
	}
	return false
}


func main(){

	file := "./callgraph.dot"

	// Parse the input file
	l,l_ := Parse(file)

	args := os.Args[1:]
	caller := args[0]
	callee := args[1]

	// TODO use l_ to map caller/callee to actual names
	fmt.Println(l_, Search(l, caller, callee))

}
