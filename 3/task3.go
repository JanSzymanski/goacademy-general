package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)
//
//Task3
//Create a program that allows the user to input a number. 
//Check whether the number lies between 1 and 10. [Variables]
//
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Is the given number between 1 and 10?")

    for scanner.Text() != "exit" {
        fmt.Print("Please provide a number: ")
        scanner.Scan()
        fmt.Println("In the loop")
        input := scanner.Text()

        converted, err := strconv.Atoi(input)
        if err != nil {
            fmt.Printf("Cannot convert %q to int\n", input)
            continue
        }
        if converted >= 1 && converted <= 10 {
            fmt.Printf("Yes, %d lays between 1 and 10.\n", converted)
        } else {
            fmt.Printf("No, %d is outside of the given range\n", converted)
        }
        
    }
    fmt.Println("Shutting down..")
    fmt.Println("Thank you for using the app")


   

}
