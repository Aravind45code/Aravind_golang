package main

import "fmt"
type passenger struct{
	Name string
	TicketNumber int
	Boarded bool
}
type Bus struct{
	FrontSeat passenger
}

func main(){
	aravind:=passenger{"Aravind",123,true}
	fmt.Println(aravind)
	//anonymus structures
	//when declared withn var give deafult but not in shorthand
	
		var bill=passenger{Name:"Bill",TicketNumber:2}
		var ella=passenger{Name:"Ella",TicketNumber:3}
	
		fmt.Println(bill,ella)
		var heidi passenger
		heidi.Name="Arun"
		heidi.TicketNumber=900
		fmt.Println(heidi)
		aravind.Boarded=false
		heidi.Boarded=true
		bus:=Bus{heidi}
		fmt.Println(bus)
		fmt.Println(bus.FrontSeat.Name)
		fmt.Println(bus.FrontSeat.Boarded)
		fmt.Println(bus.FrontSeat.TicketNumber)
	}