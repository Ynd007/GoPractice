package main

import "fmt"

func main(){
	var a [57]string
	for i:=65; i<=122; i++{
		a[i-65]=string(i)
	}
	fmt.Println(a)
}
