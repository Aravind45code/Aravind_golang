package main

import "fmt"
func Printslice(title string,slice1 []string){
	fmt.Println(title)
	for i:=0;i<len(slice1);i++{
		element:=slice1[i]
		fmt.Println(element)
	}
}
func main(){
	route:=[]string{"aravind","arun","raj","hello"}
	Printslice("route1",route)
	route=append(route,"excellent")
	Printslice("route2",route)
	fmt.Println()
	fmt.Println(route[0],"visited")
	fmt.Println(route[1],"visited")
	route=route[2:]
	fmt.Println("Remaining slice",route)

}