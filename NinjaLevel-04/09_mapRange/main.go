package main

import "fmt"

func main(){
	x:=map[string][]string{
		"bond_james":[]string{`Shaken, not stirred`, `Martinis`, `Women`},
	}
	fmt.Println(x)
for i,v:=range x{
	fmt.Println(i)
	for _,j:=range v{
		fmt.Println(j)
	}
}
}
