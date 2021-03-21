package main

import "fmt"

func main(){
	a:=10
	var b *int=&a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)

	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",&a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",*b)

}
