package main

import "fmt"

func main(){
	for i := 1; i < 10; i++{
		for j := (10 - i); j > 0; j-- {
			fmt.Print(" ");
		}
		for j := 0; j < i; j++ {
			fmt.Print(" *");
		}
		fmt.Println();
	}
}
