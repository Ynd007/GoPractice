package main

import "fmt"

func main() {
	for i := 0; i <= 30000; i++ {
		fmt.Printf("Decimal :%d\tBinary :%b\tHexa :%x\n", i, i, i)
	}
}
