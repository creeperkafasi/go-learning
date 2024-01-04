package main

import "fmt"

func main(){
	for i := 10; i > 0; i--{
		for j := i; j < 10 ; j++ {
			fmt.Print(" ");
		}
		for j := 0; j < i; j++ {
			fmt.Print(" *");
		}
		fmt.Println();
	}
}
