package main

import (
	"fmt"
	"sort"
)

func main(){
	s:=[]string{"zen","jam","ram","bheem"}
	fmt.Println(s)


	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Print(s)

}
