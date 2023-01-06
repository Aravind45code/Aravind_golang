package main

import "fmt"
type Room  struct{
	name string
	cleaned bool
} 
func checkCleanliness(rooms [4] Room){
	for i:=0;i<len(rooms);i++{
		room:=rooms[i]
		if(room.cleaned){
			fmt.Println(room.name,"is cleaned")
		}else{
			fmt.Println(room.name,"is not cleaned")
		}

	}
}
func main(){
	var myArray [3] int
	for i:=0;i<len(myArray);i++{
		fmt.Println(myArray[i])
	}
	myArray=[3]int{7,8,9}
	for i:=0;i<len(myArray);i++{
		fmt.Println(myArray[i])
	}
	newArray:=[...]int{1,2,3}
	newArray2:=[3]int{5,6,7}
	for i:=0;i<len(newArray);i++{
		fmt.Println(newArray[i])
	}
	for i:=0;i<len(newArray2);i++{
		fmt.Println(newArray2[i])
	}
	rooms:=[...]Room{
		{name:"room1"},
		{name:"room2"},
		{name:"room3"},
		{name:"room4"},}
		rooms[0].cleaned=true
	checkCleanliness(rooms)



}