package main

import "fmt"
func main(){
	//shoppingList:=make(map[string] int)
	var shoppingList map[string]int
	shoppingList["bread"]+=1
	shoppingList["milk"]=1
	shoppingList["butter"]=1
	shoppingList["rice"]=1
	fmt.Println("bread quantity",shoppingList["bread"])
	fmt.Println(shoppingList)
	delete(shoppingList,"milk")
	fmt.Println(shoppingList)
	cereal,found:=shoppingList["cereal"]
	if !found{
		fmt.Println("not found")
	}else{
		fmt.Println(cereal,"found")
	}
	totalItems:=0
	for _,amount:=range shoppingList{
		totalItems+=amount
	}
	fmt.Println(totalItems)
}