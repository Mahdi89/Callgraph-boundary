//DOT Graph analyzer
//Author: Mahdi Jelodari
//Date: April 2018
package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {

	scanstd := bufio.NewScanner(os.Stdin)

	// Get the src and dist nodes
	scanstd.Scan()
	caller := scanstd.Text()
	scanstd.Scan()
	callee := scanstd.Text()

	fmt.Println(caller, callee)
}
