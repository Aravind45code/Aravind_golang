package main

import "fmt"
func  double(x int) int{
	return x+x
}
func add(lhs,rhs int) int{
	return lhs+rhs
}
func greet(){
	fmt.Println("hello from greeting function")
}
func returnTwoNumbers() (int ,int){
	return 2,2
}
func five() int{
	return 5
}
func addThree(a,b,c int) int{
	return a+b+c
}
func greetPerson(name string){
	fmt.Println("hello",name)

}
func hiThere() string{
	return "hithere"
}
func main() {
	greet()
   dozen:=double(6)
   fmt.Println("dozen is",dozen)
   dozen2:=add(dozen,1)
   fmt.Println("the dozen2 is",dozen2)
   dozen3:=add(double(dozen2),1);
   fmt.Println("the dozen3 is",dozen3)
   a,b:=returnTwoNumbers()
   fmt.Println("the returned numbers are ",a,b)

   fmt.Println(five())
   afteraddition:=addThree(1,2,3)
   fmt.Println("The three number sum is",afteraddition)
   greetPerson("Aravind")
   fmt.Println(hiThere())


}