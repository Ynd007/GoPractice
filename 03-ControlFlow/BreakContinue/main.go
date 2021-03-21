package main

import "fmt"

func main() {
	i := 1
	count := 0

	for {

			if count==101{
				break
			} else if i%2==0{
			count++
			fmt.Println(i)
		}
		i++
	}

}