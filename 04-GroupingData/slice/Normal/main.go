package main

import "fmt"

func main(){
	slice:=[]int{1,2,3,4,5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	for i,v:=range slice{
		fmt.Println("Element at index ",i,"is ",v)
	}
}
