package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(x)
	y:=[]string{"James", "Bond", "Shaken, not stirred"}
	z:=[]string{"Miss", "Moneypenny", "Helloooooo, James."}
	a:=[][]string{y,z}
	fmt.Println(a)
}
