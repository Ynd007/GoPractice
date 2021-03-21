package main

import (
	"fmt"
	"sort"
)

type Person struct{
	Name string
	Age int
}
type ByAge []Person
func (a ByAge) Len()int 	{return len(a)}
func (a ByAge) Swap(i,j int)	{a[i],a[j]=a[j],a[i]}
func (a ByAge) Less(i,j int) bool	{return a[i].Age<a[j].Age}

func main(){
	p1:=Person{
		"krish",31,
		}
	p2:=Person{
		"mahesh",12,
	}
	person:=[]Person{p1,p2}
	sort.Sort(ByAge(person))
	fmt.Println(person)
}