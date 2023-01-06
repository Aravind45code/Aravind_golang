package main

import "fmt"
func main(){
	x:=3
switch x {
case 1:
	fmt.Println("1")
case 2:
	fmt.Println("2")
case 3:
	fmt.Println("3")
default:
	fmt.Println("other",x)
}
}
