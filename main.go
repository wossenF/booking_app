package main
import "fmt" ;

func main() {
	// fmt.Println("welcome to this booking app")
	
	// var conferenceName= "GopherCon"
	// const ticketPrice= 100
	// var remainingTickets= 70

	// fmt.Printf("hello %v conferenceName lcome to the conference \n", conferenceName)
	// fmt.Printf("The ticket price is %v and the remaining ticket is %v ", ticketPrice, remainingTickets)
	// fmt.Println("get your ticket here to attend the event")



	// var ticketCount int 
	// var totalPrice int
	// var remainingTicketsAfterPurchase int

	// ticketCount= 5
	// totalPrice= ticketCount * ticketPrice
	// remainingTicketsAfterPurchase= remainingTickets - ticketCount

	// fmt.Printf("you have purchased %T tickets and the total price is %T  and now we have %T tickets\n", ticketCount, totalPrice, remainingTicketsAfterPurchase)


	var firstName string
	var lastName string
	var email string
	var userTicket int

var availableTickets= 10


fmt.Printf("what is your name? ")
fmt.Scan(&firstName)

fmt.Printf("what is last name? ")
fmt.Scan(&lastName)

fmt.Printf("what is your email adress? ")
fmt.Scan(&email)

fmt.Printf("how many tickets do you want to purchase? ")
fmt.Scan(&userTicket)

if userTicket > availableTickets {
	fmt.Printf("sorry we only have %v tickets available", availableTickets)
	return
	
}else{
	availableTickets= availableTickets - userTicket
	fmt.Printf("Hello %v %v, welcome to the conference. you booked %v tickets, we will send the confirmation on  %v", firstName, lastName, userTicket, email)
	fmt.Printf("Now we have %v tickets available", availableTickets)
}



}