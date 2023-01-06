package main

import "fmt"
func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i;
		fmt.Println("Sum is",sum)
	}
	for sum>10{
		sum-=5
		fmt.Println("Sum is",sum)
	}
	for i:=1;i<=50;i++{
		divisibleBy3:=i%3==0
		divisibleBy5:=i%5==0
		if divisibleBy3 && divisibleBy5{
			fmt.Println("FizzBuzz")
		}else if divisibleBy3{
			fmt.Println("Fizz")
		}else{
			fmt.Println("Buzz")
		}

	}
}