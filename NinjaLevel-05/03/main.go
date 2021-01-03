package main

import "fmt"

type vehicle struct{
	doors int
	color string
}
type truck struct{
	vehicle
	fourWheel bool
}
type sedan struct{
	vehicle
	luxury bool
}
func main(){
	t:=truck{
		vehicle{
			4,
			"red",
		},
		true,
	}
	s:=sedan{
		vehicle{
			5,
			"blue",
		},
		true,
	}
	fmt.Println(t)
	fmt.Println("\t",t.doors)
	fmt.Println("\t",t.color)
	fmt.Println(s)
	fmt.Println("\t",s.doors)
	fmt.Println("\t",s.color)
}
