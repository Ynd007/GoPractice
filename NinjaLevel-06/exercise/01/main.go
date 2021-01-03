package main

import "fmt"

func EvenOdd(n int)(int,bool){
	var b bool
	var m int
	if n%2==0{
		m=1
		b=true
	}else{
		m=0
		b=false
	}
	return m,b
}
func main(){
	fmt.Println(EvenOdd(10))
	fmt.Println(EvenOdd(11))
}