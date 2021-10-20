package main

import(
	"fmt"
)
type Car struct{
	Model string
	Price int
	Name string
}

func main(){

	var a=Car{"X4",200000,"BMW"};
	fmt.Println(a);
	fmt.Println((a.Model));
	a.Model="X5";
	fmt.Println(a.Model);
}