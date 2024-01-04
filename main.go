package main

import "fmt"

func main(){
	var yükseklik int;

	fmt.Print("Üçgen Yüksekliği: ");
	_, err := fmt.Scan(&yükseklik);

	if err != nil{
		fmt.Println("Error:", err);
	}

	for i := yükseklik; i > 0; i--{
		for j := i; j < yükseklik; j++ {
			fmt.Print(" ");
		}
		for j := 0; j < i; j++ {
			fmt.Print(" ", i % 10);
		}
		fmt.Println();
	}
}
