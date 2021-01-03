package main

import "fmt"

func main(){
	sum:=SumRec(10)
	fmt.Println(sum)
}
func SumRec(n int) int{
	if n==0{
		return 0
	}
	return n+SumRec(n-1)
}

