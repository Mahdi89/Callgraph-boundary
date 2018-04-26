//DOT Graph analyzer
//Author: Mahdi Jelodari
//Date: April 2018
package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("./callgraph.dot")
	check(err)
	defer f.Close()

	scanstd := bufio.NewScanner(os.Stdin)

	// Get the src and dist nodes
	scanstd.Scan()
	caller := scanstd.Text()
	scanstd.Scan()
	callee := scanstd.Text()

	fmt.Println(caller, callee)

	// Parse the generated dot graph 
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		str := scanner.Text()
		if s.HasPrefix(str, "Node") {
			scanner.Scan()
			nxt := scanner.Text()
			if s.HasPrefix(nxt, "->") {
				scanner.Scan()
				str2 := scanner.Text()
				fmt.Println(str2)
			}
		}
	}
}
