package main

import "fmt"

func main() {
	// Variables
	age := 25
	name := "joy"
	height := 5.8562294568

	// Difference between Println vs Printf
	fmt.Println("age :", age, " height: ", height, " name: ", name)
	fmt.Println("hello world")

	fmt.Printf("age is %d",age)
	fmt.Printf(" height is %.4f\n",height)

	fmt.Printf("Type of age is %T\n",age)
	fmt.Printf("Type of name is %T\n",name)
	
	fmt.Printf("Name:%s,Age:%d,Height:%.3f\n",name,age,height)
}
