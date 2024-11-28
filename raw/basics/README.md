# Go Conference Ticket Booking Application

This is a simple **Go** console application that allows users to book tickets for a conference. The application demonstrates the use of **goroutines**, **sync.WaitGroup**, and basic user input handling in Go. It also implements validation for user input to ensure correct data entry.

---

## Features

- Allows users to book tickets for a conference.
- Validates user input for:
    - Valid first and last names.
    - A valid email address (contains `@`).
    - Valid number of tickets (not exceeding available tickets).
- Processes ticket booking in real-time.
- Sends tickets asynchronously to the user's email using a goroutine.
- Tracks and displays the first names of users who have booked tickets.

---

## Go Packages Used

1. **`fmt`**: Provides functions for formatted I/O (e.g., `Println`, `Printf`, `Scan`).
2. **`sync`**: Used for concurrency, specifically with `sync.WaitGroup` to wait for goroutines to complete.
3. **`time`**: Provides utilities for time manipulation and delays (e.g., `time.Sleep`).
4. **`basic/ticket-booking/helper`**: A custom package containing helper functions for validating user input.

---

## Variable and Function Scopes in the Application

- **Global Variables**:
    - `ConferenceName`: Stores the name of the conference (accessible globally).
    - `remainingTickets`: Tracks the number of tickets left (accessible globally).
    - `bookings`: A slice of `UserData` to store all user bookings.
    - `wg`: A `sync.WaitGroup` instance for managing goroutines.

- **Local Variables**:
    - Declared within functions, e.g., `firstName`, `lastName`, and `email` inside `getUserInput()`.
    - These variables are temporary and only exist during the function execution.

- **Functions**:
    - **Global Functions**: All functions in this program are accessible within the `main` package:
        - `main()`: Entry point of the application.
        - `greetUsers()`: Displays a welcome message.
        - `getUserInput()`: Prompts the user for input.
        - `bookTicket()`: Books tickets for the user and updates global variables.
        - `getFirstNames()`: Retrieves the first names of users who booked tickets.
        - `sendTickets()`: Simulates ticket generation and sending via email (runs in a goroutine).
        - In `map-example` directory same functionality was used with map instead of struct.  
---

## File Structure

├── main.go # Main application logic  
&nbsp;&nbsp;&nbsp;&nbsp;└── helper/  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└── validate.go # Contains the ValidateUserInput function for validating user inputs

### File Descriptions

- **`main.go`**:  
  Handles the core logic of the application, including ticket booking, user input processing, and asynchronous ticket delivery.

- **`helper/validate.go`**:  
  Contains the `ValidateUserInput` function, which validates the following:
    - Name length.
    - Presence of `@` in the email address.
    - Validity of the ticket count.

---

## How to Run the Application

### Prerequisites

- **Go** installed on your machine (version 1.20+ recommended).
- Basic understanding of running Go programs.

### Commands
1. Clone the repository:

   ```bash
   git clone https://github.com/mahfuzmission/go-ecosystem-exploration
   cd conference-ticket-booking/raw/basics
2. **Initialize and Tidy Dependencies**:
   ```bash
   go mod tidy
3. **Initialize and Tidy Dependencies**:
   ```bash
   go run .
