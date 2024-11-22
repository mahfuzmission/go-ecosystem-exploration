package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const conferenceTickets int = 50

var ConferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings [50]string // array => var bookings = [50] string{"Mahfuz", "Mission", "testUser"}
// var bookings []string // slice -> slice is array under the hood but have dynamic size
var bookings = make([]map[string]string, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTicket := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
			continue
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
			continue
		}
		if !isValidUserTicket {
			fmt.Println("number of tickets you entered is invalid")
			continue
		}

		bookTicket(userTickets, firstName, lastName, email)
		// goroutine -> 'go' keyword creates a separate thread in the background to execute the 'sendTickets' function and does the cleanup after the execution is done
		wg.Add(1)
		go sendTickets(firstName, lastName, email, userTickets)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", ConferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend the '" + ConferenceName + "'")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking) -> split the array
		//firstNames = append(firstNames, names[0])
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Map stores similar type of data
	var userDataMap = make(map[string]string)
	userDataMap["firstName"] = firstName
	userDataMap["lastName"] = lastName
	userDataMap["email"] = email
	userDataMap["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// we can use interface to pass different type of values, but it looses the strongly typed concept
	//userDataMapInterface := map-example[string]interface{}{
	//	"firstName": firstName,
	//	"lastName":  lastName,
	//	"email": email,
	//	"numberOfTickets": userTickets,
	//}

	// bookings[0] = firstName + " " + lastName // how array assigns value into specific index
	// bookings = append(bookings, firstName+" "+lastName) // how slice assigns value into next index dynamically
	bookings = append(bookings, userDataMap)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
}

func sendTickets(firstName string, lastName string, email string, userTickets uint) {
	/*
	** let consider this function takes time to complete for generating processing data, generating pdf and sending email
	 */
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf(" sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done()
}